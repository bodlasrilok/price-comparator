package main

import (
	"context"
	modelProductCrawl "github.com/tokopedia/price-comparator/internal/model/productcrawl"
	usecase "github.com/tokopedia/price-comparator/internal/usecase"
)

 //
//func getQuartileForSolds(sortedProduct []modelVasProduct.Product) model.Quartile {
//	numberOfProductPer4 := (len(sortedProduct)) / 4
//	return model.Quartile{
//		Q1: float64(sortedProduct[numberOfProductPer4].Sold.SoldCountLast30Days),
//		Q2: float64(sortedProduct[numberOfProductPer4*2].Sold.SoldCountLast30Days),
//		Q3: float64(sortedProduct[numberOfProductPer4*3].Sold.SoldCountLast30Days),
//		Q4: float64(sortedProduct[(numberOfProductPer4*4)-1].Sold.SoldCountLast30Days),
//	}
//}
//
//func calculatePrice2(products []modelVasProduct.Product, base, exponentFactor float64) (suggestedPrice []float64) {
//	//sort the products on no of items solds ascending
//	sort.Slice(products, func(i, j int) bool {
//		return products[i].Sold.SoldCountLast30Days < products[j].Sold.SoldCountLast30Days
//	})
//
//	q := getQuartileForSolds(products)
//	l := getOutlierLimit(q)
//	solds := []float64{}
//	weight := []float64{}
//	fmt.Println(l.Upper, l.Lower)
//	for _, product := range products {
//		fmt.Println("product sales are...", product.Sold.SoldCountLast30Days)
//		if float64(product.Sold.SoldCountLast30Days) < l.Upper && float64(product.Sold.SoldCountLast30Days) > l.Lower {
//			solds = append(solds, product.Price)
//			factor := math.Pow(base, (float64(len(products)) - (float64(product.Index) * exponentFactor)))
//			weight = append(weight, factor)
//		}
//	}
//	mean := stat.Mean(solds, weight)
//	std := stat.StdDev(solds, weight)
//	suggestedPrice = append(suggestedPrice, mean+(std))
//	suggestedPrice = append(suggestedPrice, mean-(std))
//	return
//
//}

func main() {
	productCrawl := modelProductCrawl.ProductCrawl{}

	//productCrawl.ProductID = 6405181
	//productCrawl.RequestID = "ps_200406184513347804"
	//productCrawl.NormalizedName = "logitech f310 gamepad"
	//productCrawl.CategoryL1ID = 2099
	//productCrawl.CategoryL2ID = 3818
	//productCrawl.CategoryL3ID = 3833
	//productCrawl.ProcessStep = 3
	//productCrawl.FilterTemplateID = 2
	//productCrawl.CreateBy = -1
	//productCrawl.Status = 1
	//productCrawl.Price = 340000
	//productCrawl.ProductURL = "https://www.tokopedia.com/shopvenus/logitech-f310-gamepad"
	//productCrawl.UpdateBy =	-1
	//productCrawl.DiscountedPrice = 0
	//productCrawl.FullCount = 3
	//productCrawl.ShopID = 200
	//productCrawl.ShopID = 0

	productCrawl.ProductID = 277881203
	productCrawl.RequestID = "ps_200406184513347804"
	productCrawl.NormalizedName = "elektrik treadmill tl 288 total gym"
	productCrawl.CategoryL1ID = 2099
	productCrawl.CategoryL2ID = 3818
	productCrawl.CategoryL3ID = 3833
	productCrawl.ProcessStep = 3
	productCrawl.FilterTemplateID = 2
	productCrawl.CreateBy = -1
	productCrawl.Status = 1
	productCrawl.Price = 5300000
	productCrawl.ProductURL = "https://www.tokopedia.com/shopvenus/logitech-f310-gamepad"
	productCrawl.UpdateBy = -1
	productCrawl.DiscountedPrice = 0
	productCrawl.FullCount = 3
	productCrawl.ShopID = 200
	productCrawl.ShopID = 0

	ctx := context.Background()

	usecase.PrintPriceSuggestions(ctx, productCrawl)

}
