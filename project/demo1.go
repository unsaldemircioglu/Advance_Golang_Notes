package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Products struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"id"`
	unitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Products
	json.Unmarshal(bodyBytes, &products)
	fmt.Println(products)
}

func AddProduct() {
	product := Product{Id: 4, ProductName: "Telephone", CategoryId: 1, UnitPrice: 6000.0}
	jsonProduct, err = json.Marshal(product)

	respons, err := http.Post("http//localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		return nil, err
	}
}
