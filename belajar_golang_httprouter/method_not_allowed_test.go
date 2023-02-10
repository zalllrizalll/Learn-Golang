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

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()
	// JIKA METHOD DARI REQUEST CLIENT SELAIN POST, JALANKAN ACTION DI BAWAH INI
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Tidak Diijinkan Mengakses")
	})
	// JIKA METHOD DARI REQUEST CLIENT ADALAH POST, JALANKAN ACTION DI BAWAH INI
	router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "POST")
	})
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Tidak Diijinkan Mengakses", string(byte))
}
