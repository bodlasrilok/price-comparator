package usecase

import (
	"context"
	"fmt"
	"github.com/tokopedia/price-comparator/internal/commons"
	"github.com/tokopedia/price-comparator/internal/model"
	modelProductCrawl "github.com/tokopedia/price-comparator/internal/model/productcrawl"
	modelVasProduct "github.com/tokopedia/price-comparator/internal/model/vasproduct"
	"github.com/tokopedia/price-comparator/internal/repository"
	"github.com/tokopedia/tdk/go/log"
	"gonum.org/v1/gonum/stat"
	"math"
	"sort"
)

func PrintPriceSuggestions(ctx context.Context, productCrawl modelProductCrawl.ProductCrawl) {
	products, _, _ := repository.GetProductsByTitle(productCrawl.NormalizedName, "TKPD277881203", 5000, 0)
	filteredProducts := products
	filteredProducts, err := repository.ClusterByCategory(products, productCrawl)
	if err != nil {
		log.Error(err, log.KV{"req_id": productCrawl.RequestID, "total products": len(products)})
	}
	if len(filteredProducts) == 0 {
		log.Infoln("filter products are zer0 for product ", productCrawl)
		repository.ForceActionDone(ctx, fmt.Sprintf("force quiting", len(products)), productCrawl)
	}
	productCrawl.CalculatedPrices = commons.GetBasicStatistic(filteredProducts)

	// calculating suggested price using formula 1
	suggestedPrice1 := calculatePrice1(filteredProducts, 2, 0.1)
	productCrawl.PriceMax = suggestedPrice1[0]
	productCrawl.PriceMin = suggestedPrice1[1]
	log.Infoln("Min and max areee....", productCrawl.PriceMin, productCrawl.PriceMax)
	log.Infoln("Suggested price using formula 1 is: ", (productCrawl.PriceMax+productCrawl.PriceMin)/2)

}

func getQuartile(sortedProduct []modelVasProduct.Product) model.Quartile {
	numberOfProductPer4 := (len(sortedProduct)) / 4
	return model.Quartile{
		Q1: sortedProduct[numberOfProductPer4].Price,
		Q2: sortedProduct[numberOfProductPer4*2].Price,
		Q3: sortedProduct[numberOfProductPer4*3].Price,
		Q4: sortedProduct[(numberOfProductPer4*4)-1].Price,
	}
}
func getOutlierLimit(q model.Quartile) model.Limit {
	return model.Limit{
		Upper: ((q.Q3 - q.Q1) * 1.5) + q.Q3,
		Lower: ((q.Q3 - q.Q1) * 1.5) - q.Q1,
	}
}
func calculatePrice1(products []modelVasProduct.Product, base, exponentFactor float64) (suggestedPrice []float64) {
	//sort the products on price ascending
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})

	q := getQuartile(products)
	l := getOutlierLimit(q)
	prices := []float64{}
	weight := []float64{}
	for _, product := range products {
		if product.Price < l.Upper && product.Price > l.Lower {
			prices = append(prices, product.Price)
			factor := math.Pow(base, (float64(len(products)) - (float64(product.Index) * exponentFactor)))
			weight = append(weight, factor)
		}
	}
	mean := stat.Mean(prices, weight)
	std := stat.StdDev(prices, weight)
	suggestedPrice = append(suggestedPrice, mean+(std))
	suggestedPrice = append(suggestedPrice, mean-(std))

	return
}
