package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"github.com/tokopedia/price-comparator/internal/model"
	modelProductCrawl "github.com/tokopedia/price-comparator/internal/model/productcrawl"
	"github.com/tokopedia/price-comparator/internal/usecase"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	inputProductHistoryFile, _ := os.Open("input/product-history.csv")
	history := csv.NewReader(bufio.NewReader(inputProductHistoryFile))
	productsData := []model.HistoryData{}

	dataMap := make(map[string][]model.HistoryData)

	inputFile, _ := os.Open("input/inputs.csv")
	reader := csv.NewReader(bufio.NewReader(inputFile))
	products := []modelProductCrawl.ProductCrawl{}

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		id, _ := strconv.ParseInt(line[1], 10, 64)
		l1, _ := strconv.ParseInt(line[3], 10, 64)
		l2, _ := strconv.ParseInt(line[4], 10, 64)
		l3, _ := strconv.ParseInt(line[5], 10, 64)
		ps, _ := strconv.Atoi(line[6])
		price, _ := strconv.ParseFloat(line[11], 64)
		shopId, _ := strconv.ParseInt(line[27], 10, 64)

		products = append(products, modelProductCrawl.ProductCrawl{
			ProductID:      id,
			CategoryL1ID:   l1,
			CategoryL2ID:   l2,
			CategoryL3ID:   l3,
			ProcessStep:    ps,
			Price:          price,
			ProductURL:     line[12],
			RequestID:      line[13],
			NormalizedName: line[15],
			ShopID:         shopId,
		})
	}

	for {
		line, error := history.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		price, _ := strconv.ParseFloat(line[1], 64)
		stock, _ := strconv.ParseInt(line[2], 10, 64)
		sold, _ := strconv.ParseInt(line[3], 10, 64)
		view, _ := strconv.ParseInt(line[4], 10, 64)

		productsData = append(productsData, model.HistoryData{
			ProductID: line[0],
			Price:     price,
			Stock:     stock,
			Sold:      sold,
			View:      view,
		})
	}

	for _, data := range productsData {
		dataMap[data.ProductID] = append(dataMap[data.ProductID], data)
	}

	ctx := context.Background()

	usecase.PrintPriceSuggestions(ctx, products[1], dataMap)

	//usecase.PrintPriceSuggestions(ctx, products[7])

	//output := []*model.Output{}
	//
	//for _, product := range products {
	//	p1, p2, p3, p4 := usecase.PrintPriceSuggestions(ctx, product, dataMap)
	//
	//	output = append(output, &model.Output{
	//		ProductID:      product.ProductID,
	//		NormalizedName: product.NormalizedName,
	//		Price:          product.Price,
	//		Price1:         p1,
	//		Price2:         p2,
	//		Price3:         p3,
	//		Price4:         p4,
	//	})
	//}
	//
	//clientsFile, err := os.OpenFile("output/results.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	//if err != nil {
	//	panic(err)
	//}
	//defer clientsFile.Close()
	//
	//if err := gocsv.UnmarshalFile(clientsFile, &output); err != nil { // Load clients from file
	//	panic(err)
	//}
	//
	//if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
	//	panic(err)
	//}
	//
	//csvContent, err := gocsv.MarshalString(&output) // Get all clients as CSV string
	//
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(csvContent)

}
