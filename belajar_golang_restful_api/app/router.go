package app

import (
	"github.com/julienschmidt/httprouter"
	"zalllrizalll/belajar_golang_restful_api/controller"
	"zalllrizalll/belajar_golang_restful_api/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	// GET ALL Data
	router.GET("/api/categories", categoryController.FindAll)
	// GET Data by ID
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	// POST Data
	router.POST("/api/categories", categoryController.Create)
	// Update Data
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	// Delete Data
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// Setting ErrorHandler => berfungsi untuk menghandle sebuah pesan error untuk disampaikan ke user atau client
	router.PanicHandler = exception.ErrorHandler
	return router
}
