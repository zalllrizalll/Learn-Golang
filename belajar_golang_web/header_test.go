package belajargolangweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type") // Mengambil suatu values data berdasarkan key string
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json") // application/json merupakan values dari key string content-type
	recorder := httptest.NewRecorder()
	RequestHeader(recorder, request)
	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Powered-By", "Males Ngoding")
	fmt.Fprint(writer, "Berhasil!")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil) // values request dari client untuk diteruskan ke server
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	ResponseHeader(recorder, request)
	response := recorder.Result() // values respon server ke client
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Header.Get("powered-by"))
}
