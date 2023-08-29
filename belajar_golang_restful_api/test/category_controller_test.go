package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
	"zalllrizalll/belajar_golang_restful_api/app"
	"zalllrizalll/belajar_golang_restful_api/controller"
	"zalllrizalll/belajar_golang_restful_api/helper"
	"zalllrizalll/belajar_golang_restful_api/middleware"
	"zalllrizalll/belajar_golang_restful_api/model/domain"
	"zalllrizalll/belajar_golang_restful_api/repository"
	"zalllrizalll/belajar_golang_restful_api/service"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	// Koneksi database
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	// Router
	router := app.NewRouter(categoryController)
	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE CATEGORY")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	// REQUEST : PERMINTAAN CLIENT UNTUK DITERUSKAN KE SERVER
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	// RECORDER : RESPONS DARI SERVER UNTUK DIKEMBALIKAN DATA YANG DIMINTA DARI CLIENT
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// RESPONS SERVER PADA CLIENT
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// JSON UNMARSHAL => MENGKONSVERSI DATA JSON KE GOLANG
	err := json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)
	// COBA TEST DATA
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name" : ""}`)
	// REQUEST : PERMINTAAN CLIENT UNTUK DITERUSKAN KE SERVER
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	// RECORDER : RESPONS DARI SERVER UNTUK DIKEMBALIKAN DATA YANG DIMINTA DARI CLIENT
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// RESPONS SERVER PADA CLIENT
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// JSON UNMARSHAL => MENGKONVERSI DATA JSON KE GOLANG
	err := json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)
	// COBA TEST DATA
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	// CREATE TX *SQL.TX
	tx, _ := db.Begin()
	// CREATE DATA TERLEBIH DAHULU
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	// COMMIT TX *SQL.TX
	tx.Commit()
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	// REQUEST : PERMINTAAN CLIENT UNTUK DITERUSKAN KE SERVER
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	// RECORDER : RESPONS DARI SERVER UNTUK DIKEMBALIKAN DATA YANG DIMINTA DARI CLIENT
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// RESPONS SERVER PADA CLIENT
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// JSON UNMARSHAL => MENGKONSVERSI DATA JSON KE GOLANG
	err := json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)
	// COBA TEST DATA
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	// CREATE TX *SQL.TX
	tx, _ := db.Begin()
	// CREATE DATA TERLEBIH DAHULU
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	// COMMIT TX *SQL.TX
	tx.Commit()
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name" : ""}`)
	// REQUEST : PERMINTAAN CLIENT UNTUK DITERUSKAN KE SERVER
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	// RECORDER : RESPONS DARI SERVER UNTUK DIKEMBALIKAN DATA YANG DIMINTA DARI CLIENT
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// RESPONS SERVER PADA CLIENT
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// JSON UNMARSHAL => MENGKONVERSI DATA JSON KE GOLANG
	err := json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)
	// COBA TEST DATA
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	// CREATE TX *SQL.TX
	tx, _ := db.Begin()
	// CREATE DATA TERLEBIH DAHULU
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	// COMMIT TX *SQL.TX
	tx.Commit()
	router := setupRouter(db)
	// REQUEST : PERMINTAAN CLIENT UNTUK DITERUSKAN KE SERVER
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	// RECORDER : RESPONS DARI SERVER UNTUK DIKEMBALIKAN DATA YANG DIMINTA DARI CLIENT
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// RESPONS SERVER PADA CLIENT
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// JSON UNMARSHAL => MENGKONVERSI DATA JSON KE GOLANG
	err := json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)
	// COBA TEST DATA
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)
	// REQUEST : PERMINTAAN CLIENT UNTUK DITERUSKAN KE SERVER
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	// RECORDER : RESPONS DARI SERVER UNTUK DIKEMBALIKAN DATA YANG DIMINTA DARI CLIENT
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// RESPONS SERVER PADA CLIENT
	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// JSON UNMARSHAL => MENGKONVERSI DATA JSON KE GOLANG
	err := json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)
	// COBA TEST DATA
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	// CREATE TX *SQL.TX
	tx, _ := db.Begin()
	// CREATE DATA TERLEBIH DAHULU
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	// COMMIT TX *SQL.TX
	tx.Commit()
	router := setupRouter(db)
	// REQUEST : PERMINTAAN CLIENT UNTUK DITERUSKAN KE SERVER
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	// RECORDER : RESPONS DARI SERVER UNTUK DIKEMBALIKAN DATA YANG DIMINTA DARI CLIENT
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// RESPONS SERVER PADA CLIENT
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// JSON UNMARSHAL => MENGKONSVERSI DATA JSON KE GOLANG
	err := json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)
	// COBA TEST DATA
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)
	// REQUEST : PERMINTAAN CLIENT UNTUK DITERUSKAN KE SERVER
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	// RECORDER : RESPONS DARI SERVER UNTUK DIKEMBALIKAN DATA YANG DIMINTA DARI CLIENT
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// RESPONS SERVER PADA CLIENT
	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// JSON UNMARSHAL => MENGKONSVERSI DATA JSON KE GOLANG
	err := json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)
	// COBA TEST DATA
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListCategoriesSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	// CREATE TX *SQL.TX
	tx, _ := db.Begin()
	// CREATE DATA TERLEBIH DAHULU
	categoryRepository := repository.NewCategoryRepository()
	// INPUT DATA DARI CLIEN
	category1 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	category2 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Komputer",
	})
	// COMMIT TX *SQL.TX
	tx.Commit()
	router := setupRouter(db)
	// REQUEST : PERMINTAAN CLIENT UNTUK DITERUSKAN KE SERVER
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	// RECORDER : RESPONS DARI SERVER UNTUK DIKEMBALIKAN DATA YANG DIMINTA DARI CLIENT
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// RESPONS SERVER PADA CLIENT
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// JSON UNMARSHAL => MENGKONVERSI DATA JSON KE GOLANG
	err := json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)
	// COBA TEST DATA
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var categories = responseBody["data"].([]interface{})
	// DATA 1
	categoryResponse1 := categories[0].(map[string]interface{})
	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Name, categoryResponse1["name"])
	// DATA 2
	categoryResponse2 := categories[1].(map[string]interface{})
	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoryResponse2["name"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)
	// REQUEST : PERMINTAAN CLIENT UNTUK DITERUSKAN KE SERVER
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "SALAH")

	// RECORDER : RESPONS DARI SERVER UNTUK DIKEMBALIKAN DATA YANG DIMINTA DARI CLIENT
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// RESPONS SERVER PADA CLIENT
	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// JSON UNMARSHAL => MENGKONVERSI DATA JSON KE GOLANG
	err := json.Unmarshal(body, &responseBody)
	helper.PanicIfError(err)
	// COBA TEST DATA
	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
