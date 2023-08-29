package belajargolangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, err := os.Create("CustomerOut.json")
	if err != nil {
		panic(err)
	}
	// Panggil si Library json.NewEncoder()
	encoder := json.NewEncoder(writer)
	customer := &Customer{
		FirstName:  "Muhammad",
		MiddleName: "Rizal",
		LastName:   "Pratama",
		Age:        20,
		Married:    false,
		Hobbies: []string{
			"Gaming", "Reading", "Coding",
		},
		Addresses: []Address{
			{
				Street:     "Jalan Pahlawan",
				Country:    "Indonesian",
				PostalCode: "50149",
			},
		},
	}
	err = encoder.Encode(customer)
	if err != nil {
		panic(err)
	}
}
