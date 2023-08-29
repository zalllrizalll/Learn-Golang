package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArrayEncode(t *testing.T) {
	customer := &Customer{
		FirstName:  "Muhammad",
		MiddleName: "Rizal",
		LastName:   "Pratama",
		Hobbies: []string{
			"Gaming",
			"Reading",
			"Coding",
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Muhammad","MiddleName":"Rizal","LastName":"Pratama","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	// KONVERSI FILE JSON TIPE DATA INTERFACE{} MENJADI []BYTE
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Muhammad",
		MiddleName: "Rizal",
		LastName:   "Pratama",
		Age:        20,
		Married:    false,
		Hobbies: []string{
			"Gaming",
			"Reading",
			"Coding",
		},
		Addresses: []Address{
			{
				Street:     "Jalan Pahlawan",
				Country:    "Indonesia",
				PostalCode: "50149",
			},
			{
				Street:     "Jalanin aja dulu",
				Country:    "Maroko",
				PostalCode: "75839",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Muhammad","MiddleName":"Rizal","LastName":"Pratama","Age":20,"Married":false,"Hobbies":["Gaming","Reading","Coding"],"Addresses":[{"Street":"Jalan Pahlawan","Country":"Indonesia","PostalCode":"50149"},{"Street":"Jalanin aja dulu","Country":"Maroko","PostalCode":"75839"}]}`
	// KONVERSI FILE JSON TIPE DATA INTERFACE{} MENJADI []BYTES
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.Addresses)
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Pahlawan","Country":"Indonesia","PostalCode":"50149"},{"Street":"Jalanin aja dulu","Country":"Maroko","PostalCode":"75839"}]`
	// KONVERSI FILE JSON TIPE DATA INTERFACE{} MENJADI []BYTES
	jsonBytes := []byte(jsonString)

	Addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, Addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(Addresses)
}

func TestOnlyJsonArrayComplex(t *testing.T) {
	Addresses := []Address{
		{
			Street:     "Jalan Pahlawan",
			Country:    "Indonesia",
			PostalCode: "50149",
		},
		{
			Street:     "Jalanin aja dulu",
			Country:    "Maroko",
			PostalCode: "75839",
		},
	}

	bytes, err := json.Marshal(Addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
