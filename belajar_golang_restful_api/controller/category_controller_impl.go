package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"zalllrizalll/belajar_golang_restful_api/helper"
	"zalllrizalll/belajar_golang_restful_api/model/web"
	"zalllrizalll/belajar_golang_restful_api/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)
	categoryId := params.ByName("categoryId")
	// Konversi string => integer
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id
	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	categoryId := params.ByName("categoryId")
	// Konversi string => integer
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	categoryId := params.ByName("categoryId")
	// Konversi string => integer
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.Response{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
