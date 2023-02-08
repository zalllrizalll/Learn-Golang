package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleWare *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("Before Execute Handler")
	middleWare.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (handler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("Recover : ", err)
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	handler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware!")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(writer, "Hello Foo!")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		panic("Ups")
	})

	logMiddleWare := LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: &logMiddleWare,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}
	/**
	Langkah-Langkah Data dari Client => Server dieksekusi
	1. Client mengirim request ke server
	2. Ditangkap pesan request tersebut oleh server
	3. Server meneruskan pesan request tersebut ke errorHandler
	4. errorHandler menangkap pesan request tersebut dan diteruskan ke logMiddleWare, Apabila ditengah perjalanan tersebut terjadi error(panic) maka si errorHandler akan menjalankan si defer func() tersebut
	5. logMiddleWare menangkap pesan dari errorHandler dan diteruskan ke mux
	6. Mux akan mengeksekusi pesan request tersebut dan akan mengembalikan response server ke Client
	*/

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
