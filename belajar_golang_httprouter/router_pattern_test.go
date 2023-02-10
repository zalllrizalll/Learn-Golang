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

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Product " + id + " item " + itemId
		fmt.Fprint(writer, text)
	})
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/product/1/items/254", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)
	// PENGECEKAN VALUE => APAKAH (OBJECT 1 == OBJECT 2)
	assert.Equal(t, "Product 1 item 254", string(byte))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		text := "Image : " + image
		fmt.Fprint(writer, text)
	})
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/images/logo/profile.jpeg", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)
	// PENGECEKAN VALUE => APAKAH (OBJECT 1 == OBJECT 2)
	assert.Equal(t, "Image : /logo/profile.jpeg", string(byte))
}
