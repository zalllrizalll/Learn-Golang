package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MENAMBAH DATA COOKIE
func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)                     // REQUEST DARI CLIENT
	cookie.Name = "X-RZL-Name"                     // KEY
	cookie.Value = request.URL.Query().Get("name") // VALUE
	cookie.Path = "/"                              // DATA COOKIE TERSEBUT DAPAT DIAKSES DI HALAMAN APA SAJA

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success create cookie!")
}

// MENGAMBIL DATA COOKIE
func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-RZL-Name") // MENGAMBIL VALUE COOKIE BERDASARKAN KEY
	if err != nil {
		fmt.Fprint(writer, "Not cookie found")
	} else {
		fmt.Fprintf(writer, "Success login as %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()                // METHOD PENAMPUNG DARI BEBERAPA HANDLER
	mux.HandleFunc("/set-cookie", SetCookie) // HANDLER SET-COOKIE
	mux.HandleFunc("/get-cookie", GetCookie) // HANDLER GET-COOKIE

	server := http.Server{
		Addr:    "localhost:8080", // SERVER LOCAL UNTUK MENJALANKAN CODE PROGRAM
		Handler: mux,              // SET HANDLER
	}

	err := server.ListenAndServe() // METHOD KONEKSI KE SERVER
	if err != nil {
		panic(err)
	}
}

// MENGIRIM DATA KE WEB BROWSER UNTUK DISIMPAN KE DALAM COOKIES
func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/?name=Farah", nil)
	recorder := httptest.NewRecorder()

	// PANGGIL FUNC SET-COOKIE
	SetCookie(recorder, request)
	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies { // FOR EACH UNTUK LOOPING DATA DARI ARRAY
		fmt.Printf("cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

// MENDAPATKAN DATA COOKIES YANG DISIMPAN DI WEB BROWSER
func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie) // MEMBUAT OBJECT COOKIE
	cookie.Name = "X-RZL-Name"
	cookie.Value = "Farah"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	// PANGGIL FUNC GET-COOKIE
	GetCookie(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
