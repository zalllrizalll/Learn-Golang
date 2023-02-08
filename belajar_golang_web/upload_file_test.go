package belajargolangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	err := myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	//request.ParseMultipartForm(32 << 20) // Default => server hanya dapat menampung file sebesar 32 mb
	// Ambil file input dari client
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	// Mengarahkan lokasi file input dari client dimana file tersebut akan disimpan
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	// Simpan file client yang telah diarahkan ke fileDestination
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := request.FormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

//go:embed resources/Yt-Banner.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	// variable penampung file yang dikirim oleh client
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	// Method WriteField => digunakan untuk upload file berupa input text
	writer.WriteField("name", "Rizal Pratama")
	file, _ := writer.CreateFormFile("file", "YoutubeBanner.png")
	file.Write(uploadFileTest)
	// Hentikan eksekusi program agar tidak terjadi memori yang terus berjalan di latar belakang
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	// Panggil Handler
	Upload(recorder, request)
	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}

func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
