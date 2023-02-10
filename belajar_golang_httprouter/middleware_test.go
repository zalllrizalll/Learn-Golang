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

type LogMiddleWare struct {
	http.Handler
}

func (middleware *LogMiddleWare) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Receive Request")
	middleware.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Middleware!")
	})

	middleWare := LogMiddleWare{
		Handler: router,
	}
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	middleWare.ServeHTTP(recorder, request)
	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Middleware!", string(byte))
}
