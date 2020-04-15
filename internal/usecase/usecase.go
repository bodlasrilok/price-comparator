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

func PrintPriceSuggestions(ctx context.Context, productCrawl modelProductCrawl.ProductCrawl, data map[string][]model.HistoryData) (price1, price2, price3 float64) {
	products, _, _ := repository.GetProductsByTitle(ctx, productCrawl.NormalizedName, "TKPD277881203", 5000, 0)
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
	//
	//var ran int
	//if len(filteredProducts) > 14 {
	//	ran = 14
	//} else {
	//	ran = len(filteredProducts)
	//}

	// calculating suggested price using formula 1
	suggestedPrice1 := calculatePrice1(filteredProducts, 2, 0.1)
	productCrawl.PriceMax = suggestedPrice1[0]
	productCrawl.PriceMin = suggestedPrice1[1]
	price1 = (productCrawl.PriceMax + productCrawl.PriceMin) / 2
	log.Infoln("Min and max areee....", productCrawl.PriceMin, productCrawl.PriceMax)
	log.Infoln("Suggested price using formula 1 is: ", (productCrawl.PriceMax+productCrawl.PriceMin)/2, "for products id is: ", productCrawl.ProductID)

	// calculating suggested price using formula 2
	suggestedPrice2 := calculatePrice2(filteredProducts, 2, 0.1)
	productCrawl.PriceMax = suggestedPrice2[0]
	productCrawl.PriceMin = suggestedPrice2[1]
	price2 = (productCrawl.PriceMax + productCrawl.PriceMin) / 2
	log.Infoln("Min and max areee....", productCrawl.PriceMin, productCrawl.PriceMax)
	log.Infoln("Suggested price using formula 2 is: ", (productCrawl.PriceMax+productCrawl.PriceMin)/2, "for products id is: ", productCrawl.ProductID)

	// calculating suggested price using formula 3
	suggestedPrice3 := calculatePrice3(filteredProducts, 2, 0.1)
	productCrawl.PriceMax = suggestedPrice3[0]
	productCrawl.PriceMin = suggestedPrice3[1]
	price3 = (productCrawl.PriceMax + productCrawl.PriceMin) / 2
	log.Infoln("Min and max areee....", productCrawl.PriceMin, productCrawl.PriceMax)
	log.Infoln("Suggested price using formula 3 is: ", (productCrawl.PriceMax+productCrawl.PriceMin)/2, "for products id is: ", productCrawl.ProductID)

	//calculating suggested price using formula 4
	suggestedPrice4 := calculatePrice4(filteredProducts, 2, 0.1, data)
	productCrawl.PriceMax = suggestedPrice4[0]
	productCrawl.PriceMin = suggestedPrice4[1]
	log.Infoln("Min and max areee....", productCrawl.PriceMin, productCrawl.PriceMax)
	log.Infoln("Suggested price using formula 4 is: ", (productCrawl.PriceMax+productCrawl.PriceMin)/2, "for products id is: ", productCrawl.ProductID)
	return
}

func getQuartile(sortedProduct []modelVasProduct.Product) model.Quartile {
	numberOfProductPer4 := (len(sortedProduct)) / 4
	return model.Quartile{
		Q1: sortedProduct[numberOfProductPer4-1].Price,
		Q2: sortedProduct[numberOfProductPer4*2-1].Price,
		Q3: sortedProduct[numberOfProductPer4*3-1].Price,
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
	if len(products) < 4 {
		suggestedPrice = append(suggestedPrice, 0)
		suggestedPrice = append(suggestedPrice, 0)

		return
	}
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
	if len(prices) > 0 {

		mean := stat.Mean(prices, weight)
		std := stat.StdDev(prices, weight)
		suggestedPrice = append(suggestedPrice, mean+(std))
		suggestedPrice = append(suggestedPrice, mean-(std))
	} else {
		suggestedPrice = append(suggestedPrice, 0)
		suggestedPrice = append(suggestedPrice, 0)
	}

	return
}
func calculatePrice2(products []modelVasProduct.Product, base, exponentFactor float64) (suggestedPrice []float64) {
	if len(products) < 4 {
		suggestedPrice = append(suggestedPrice, 0)
		suggestedPrice = append(suggestedPrice, 0)

		return
	}
	//sort the products on price ascending
	sort.Slice(products, func(i, j int) bool {
		return products[i].Sold.SoldCountLast30Days < products[j].Sold.SoldCountLast30Days
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
	if len(prices) > 0 {

		mean := stat.Mean(prices, weight)
		std := stat.StdDev(prices, weight)
		suggestedPrice = append(suggestedPrice, mean+(std))
		suggestedPrice = append(suggestedPrice, mean-(std))
	} else {
		suggestedPrice = append(suggestedPrice, 0)
		suggestedPrice = append(suggestedPrice, 0)
	}

	return
}
func calculatePrice3(products []modelVasProduct.Product, base, exponentFactor float64) (suggestedPrice []float64) {

	//sort the products on price ascending
	sort.Slice(products, func(i, j int) bool {
		return products[i].Sold.SoldCountLast30Days < products[j].Sold.SoldCountLast30Days
	})

	prices := []float64{}
	weight := []float64{}
	var num float64
	var den float64

	count := 0

	for count < len(products) && products[count].Sold.SoldCountLast30Days < 0 {
		prices = append(prices, products[count].Price)
		count = count + 1
	}
	den = den + float64(count)

	count = count + 1

	var sp1 []float64
	if count < len(products) {
		sp1 = calculatePrice1(products[:count], base, exponentFactor)
	} else {
		sp1 = calculatePrice1(products, base, exponentFactor)
	}

	num = num + den*(sp1[0]+sp1[1])/2

	fmt.Println((sp1[0] + sp1[1]) / 2)

	for count < len(products) {
		prices = append(prices, products[count].Price)
		factor := math.Pow(base, (float64(len(products)) - (float64(products[count].Index) * exponentFactor)))
		weight = append(weight, factor)
		num = num + products[count].Price*float64(products[count].Sold.SoldCountLast30Days)
		den = den + float64(products[count].Sold.SoldCountLast30Days)
		count = count + 1
	}

	if den == 0 {
		if len(prices) > 0 {

			mean := stat.Mean(prices, weight)
			std := stat.StdDev(prices, weight)
			suggestedPrice = append(suggestedPrice, mean+(std))
			suggestedPrice = append(suggestedPrice, mean-(std))
		} else {
			suggestedPrice = append(suggestedPrice, 0)
			suggestedPrice = append(suggestedPrice, 0)
		}
	} else {
		suggestedPrice = append(suggestedPrice, num/den)
		suggestedPrice = append(suggestedPrice, num/den)
	}

	return
}

func calculatePrice4(products []modelVasProduct.Product, base, exponentFactor float64, data map[string][]model.HistoryData) (suggestedPrice []float64) {

	fmt.Println("datamap is ", data["TKPD710539777"])

	//Carries list of products from history
	productsList := []model.HistoryData{}
	for _, product := range products {
		for _, dataChunk := range data[product.ProductID] {
			if dataChunk.View < 0 || dataChunk.Sold < 0 {
				dataChunk.CNI = -1
			} else if dataChunk.View == 0 || dataChunk.Sold == 0 {
				dataChunk.CNI = 0
			} else {
				dataChunk.CNI = float64(dataChunk.Sold) / float64(dataChunk.View)
			}
			productsList = append(productsList, dataChunk)
		}
	}

	//sort the products on price ascending
	sort.Slice(productsList, func(i, j int) bool {
		return productsList[i].CNI < productsList[j].CNI
	})

	count := 0
	prices := []float64{}
	var num float64
	var den float64

	for count < len(productsList) && productsList[count].CNI < 0 {
		prices = append(prices, productsList[count].Price)
		num = num + productsList[count].Price
		count = count + 1
	}

	num = num * (float64(count) / 10) / float64(count)
	den = den + (float64(count)/10)/float64(count)
	count = count + 1
	for count < len(productsList) {
		num = num + productsList[count].Price * productsList[count].CNI
		den = den + productsList[count].CNI
	}

	if den == 0 {
		suggestedPrice = append(suggestedPrice, 0)
		suggestedPrice = append(suggestedPrice, 0)
	} else {
		suggestedPrice = append(suggestedPrice, num/den)
		suggestedPrice = append(suggestedPrice, num/den)
	}

	return
}
