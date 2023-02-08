package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello, world!")
	}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	// Home page
	mux.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Selamat Datang di Golang Web!")
	})
	// Products Page
	mux.HandleFunc("/products", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Ini Halaman Products")
	})
	// About Page
	mux.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Ini Halaman About")
	})
	// Koneksi ke localhost
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, request.Method)
		fmt.Fprint(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
