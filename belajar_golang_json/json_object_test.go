package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJsonObject(t *testing.T) {
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
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
