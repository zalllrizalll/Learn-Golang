package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello Rizal Pratama , I will talking about the HttpRouter.", string(byte))
}

func TestServeFileGoodbye(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/files/goodbye.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	byte, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Goodbye HttpRouter", string(byte))
}
