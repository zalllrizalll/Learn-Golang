package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("Function").Parse(`{{.SayHello "Farah"}}`))
	t.ExecuteTemplate(writer, "Function", MyPage{
		Name: "Rizal Pratama",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	// Panggil Handler
	TemplateFunction(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("Function").Parse(`{{ len .Name}}`))
	t.ExecuteTemplate(writer, "Function", MyPage{
		Name: "Rizal Pratama",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	// Panggil Handler
	TemplateFunctionGlobal(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	// Create Global Function
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"uppercase": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	// Parse Data
	t = template.Must(t.Parse(`{{ uppercase .Name}}`))
	// Eksekusi Program
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Rizal Pratama",
	})

}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	// Panggil Handler
	TemplateFunctionCreateGlobal(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

func TemplateFunctionCreateGlobalPipelines(writer http.ResponseWriter, request *http.Request) {
	// Create Global Function
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"uppercase": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	// Parse Data
	t = template.Must(t.Parse(`{{ sayHello .Name | uppercase}}`))
	// Eksekusi Program
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Rizal Pratama",
	})

}

func TestTemplateFunctionCreateGlobalPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	// Panggil Handler
	TemplateFunctionCreateGlobalPipelines(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
