package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"zalllrizalll/belajar_golang_restful_api/app"
	"zalllrizalll/belajar_golang_restful_api/controller"
	"zalllrizalll/belajar_golang_restful_api/helper"
	"zalllrizalll/belajar_golang_restful_api/middleware"
	"zalllrizalll/belajar_golang_restful_api/repository"
	"zalllrizalll/belajar_golang_restful_api/service"
)

func main() {
	// KONEKSI KE DATABASE belajar_golang_restful_api
	db := app.NewDB()
	// VALIDASI DATA
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	// Setting Server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	errorServer := server.ListenAndServe()
	helper.PanicIfError(errorServer)
}
