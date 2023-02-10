package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		text := "Product " + id
		fmt.Fprint(writer, text)
	})
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/product/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)
	// PENGECEKAN VALUE => APAKAH (OBJECT 1 == OBJECT 2)
	assert.Equal(t, "Product 1", string(byte))
}
