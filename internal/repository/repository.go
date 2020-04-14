package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/montanaflynn/stats"
	modelProductCrawl "github.com/tokopedia/price-comparator/internal/model/productcrawl"
	modelVasProduct "github.com/tokopedia/price-comparator/internal/model/vasproduct"
	"github.com/tokopedia/tdk/go/log"
	"github.com/tokopedia/tdk/go/tracer"
	"gopkg.in/olivere/elastic.v6"
	"math"
)

func GetProductsByTitle(productTitle, excludeProductID string, limit, offset int) (result []modelVasProduct.Product, total int64, err error) {
	ctx := context.Background()
	esClient, err := elastic.NewClient(elastic.SetURL("http://172.21.248.170:9200"))
	if err != nil {
		fmt.Println("error is ", err)
	}
	boolQuery := elastic.NewBoolQuery().Must(
		elastic.NewQueryStringQuery(productTitle).Field("normalized_name").DefaultOperator("AND"),
		elastic.NewTermQuery("active", true),
	)

	if excludeProductID != "" {
		boolQuery.MustNot(elastic.NewIdsQuery("_doc").Ids(excludeProductID))
	}

	res, err := esClient.Search().Index("product-vas").Size(limit).From(offset).Query(boolQuery).Sort("_id", true).Do(ctx)
	if err != nil {
		log.ErrorWithFields("error can't search product from elasticsearch: "+err.Error(), log.KV{"productTitle": productTitle})
		return
	}

	total = res.Hits.TotalHits

	for _, pr := range res.Hits.Hits {

		product := modelVasProduct.Product{}
		if errUnmarshal := json.Unmarshal(*pr.Source, &product); errUnmarshal == nil {

			result = append(result, product)
		}

	}

	return
}

func ClusterByCategory(list []modelVasProduct.Product, product modelProductCrawl.ProductCrawl) (relevantList []modelVasProduct.Product, err error) {

	min, max := getMinMaxPriceByThreshold(product.Price, 2*50)
	log.Infoln("min ", min, "max ", max, "product price ", product.Price)

	catIDs := make(map[int64]string)
	for _, pr := range list {
		catIDs[pr.CategoryL3ID] = pr.Name
	}
	log.Infoln("list of cat id are : ", catIDs, "length of cat ids : ", len(catIDs))

	minDiff := product.Price
	var relevantCat []int64

	for k := range catIDs {

		prices := []float64{}
		count := 0
		for _, pr := range list {
			if pr.CategoryL3ID == k {
				count = count + 1
				fmt.Println("count is : ", count, "for cat id : ", k)
				prices = append(prices, pr.Price)
			}
		}

		avgPrices, _ := stats.Mean(prices)
		diff := math.Abs(avgPrices - product.Price)
		log.Infoln("averange prices ", avgPrices, "differnce in math abs ", diff, "mindifff", float64(minDiff), "for cat id : ", k)
		if avgPrices >= float64(min) && avgPrices <= float64(max) && diff <= minDiff {
			log.Infoln("In success condition for cat id ,", k)
			minDiff = diff
			relevantCat = append(relevantCat, k)
		}

	}

	log.Infoln("relevant cat id are : ", relevantCat, "list is :", relevantCat)

	for _, pr := range list {
		for _, k := range relevantCat {
			if pr.CategoryL3ID == k {
				relevantList = append(relevantList, pr)
			}
		}
	}

	if len(relevantList) == 0 {
		err = errors.New("no relevant category")
		log.Infoln("product ->", product.ProductID, " expected category ->", relevantCat, " count ->", len(relevantList), " input data ->", len(list),
			" clustered category ->", catIDs)
		return
	}

	return
}
func getMinMaxPriceByThreshold(price, threshold float64) (min, max int64) {
	if threshold <= 0 {
		return
	}

	min = int64(price - (price * threshold / 100))
	max = int64(price + (price * threshold / 100))

	return
}

func ForceActionDone(ctx context.Context, reason string, productCrawl modelProductCrawl.ProductCrawl) {
	productCrawl.ProcessStep = 3
	productCrawl.CalculatedPrices.Messages = reason
	errUpdate := update(ctx, productCrawl)
	if errUpdate != nil {
		log.Error("force-update -> error: " + errUpdate.Error())
	}
}

func update(ctx context.Context, data modelProductCrawl.ProductCrawl) (err error) {
	span, ctx := tracer.StartFromContext(ctx)
	defer span.Finish()
	fmt.Println("updating values in db is done")

	return
}
