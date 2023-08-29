package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Prices   int64  `json:"prices"`
	ImageURL string `json:"image_url"`
}

func TestJsonTagEncode(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Pepsodent",
		Prices:   15000,
		ImageURL: "http://example.com/pepsodent.png",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Pepsodent","prices":15000,"image_url":"http://example.com/pepsodent.png"}`
	jsonBytes := []byte(jsonString)

	products := &Product{}
	err := json.Unmarshal(jsonBytes, products)
	if err != nil {
		panic(err)
	}
	fmt.Println(products)
}
