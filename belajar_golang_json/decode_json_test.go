package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"FirstName" : "Muhammad", "MiddleName" : "Rizal", "LastName" : "Pratama", "Age" : 20, "Married" : false}`
	// KONVERSI FILE JSON TIPE DATA STRING => MENJADI KE BENTUK SLICE OF BYTE
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	// JSON DECODE => MENGKONVERSI SUATU DATA DARI FILE JSON MENJADI FILE GOLANG
	// DENGAN MENGGUNAKAN SEBUAH LIBRARY PACKAGE JSON DI DALAM SEBUAH BAHASA PEMROGRAMAN GOLANG
	// DI DALAM LIBRARY TERSEBUT TERDAPAT SEBUAH METHOD YAITU JSON.UNMARSHAL(INTERFACE{}) YANG BERFUNGSI MENGKONVERSI DATA FILE JSON KE GOLANG
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
