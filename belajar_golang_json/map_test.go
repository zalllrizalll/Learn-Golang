package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id": "P001", "name": "Pepsodent", "prices": 15000, "image_url": "http://example.com/pepsodent.png"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["name"])
	fmt.Println(result["prices"])
}

func TestMapEncode(t *testing.T) {
	products := map[string]interface{}{
		"id":        "P001",
		"name":      "Pepsodent",
		"prices":    15000,
		"image_url": "http://example.com/pepsodent.png",
	}

	bytes, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
