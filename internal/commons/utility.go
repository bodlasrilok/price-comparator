package commons

import (
	"github.com/montanaflynn/stats"
	modelProductCrawl "github.com/tokopedia/price-comparator/internal/model/productcrawl"
	modelVasProduct "github.com/tokopedia/price-comparator/internal/model/vasproduct"
)

func GetBasicStatistic(filteredProducts []modelVasProduct.Product) (result modelProductCrawl.CalculatedPrices) {
	if len(filteredProducts) == 0 {
		return
	}

	prices := []float64{}

	for _, fp := range filteredProducts {
		prices = append(prices, fp.Price)
	}

	result.PriceCount = len(filteredProducts)
	result.PriceMeans, _ = stats.Mean(prices)
	result.PriceMedian, _ = stats.Median(prices)
	result.PriceMax, _ = stats.Max(prices)
	result.PriceMin, _ = stats.Min(prices)

	return
}
