package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Product struct {
	Data        string  `json:"Data"`
	ProductName string  `json:"Produto Coletado"`
	Price       float32 `json:"Preço Coletado"`
	LinkUrl     string  `json:"Link URL"`
}

func main() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	data_lines := strings.Split(string(data), "\n")

	var products []Product
	var product Product

	for _, line := range data_lines {
		if line == "" {
			continue
		}
		err := json.Unmarshal([]byte(line), &product)
		if err != nil {
			fmt.Println("Error unmarshalling:", err)
			return
		}
		products = append(products, product)
	}
	fmt.Println("Total products:", len(products))
}
