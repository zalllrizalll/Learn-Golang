package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// JSON ENCODE => MENGKONVERSI SUATU DATA DARI FILE GOLANG MENJADI FILE JSON
// DENGAN MENGGUNAKAN SEBUAH LIBRARY PACKAGE JSON DI DALAM SEBUAH BAHASA PEMROGRAMAN GOLANG
// DI DALAM LIBRARY TERSEBUT TERDAPAT SEBUAH METHOD YAITU JSON.MARSHAL(INTERFACE{}) YANG BERFUNGSI MENGKONVERSI DATA FILE GOLANG KE JSON
func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
func TestEncodeJson(t *testing.T) {
	logJson("Rizal")
	logJson(1)
	logJson(true)
	logJson([]string{"Muhammad", "Rizal", "Pratama"})
}
