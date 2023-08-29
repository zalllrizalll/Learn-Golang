package exception

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"zalllrizalll/belajar_golang_restful_api/helper"
	"zalllrizalll/belajar_golang_restful_api/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, error interface{}) {

	if notFoundError(writer, request, error) {
		return
	}

	if validationErrors(writer, request, error) {
		return
	}
	internalServerError(writer, request, error)
}
func internalServerError(writer http.ResponseWriter, request *http.Request, error interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.Response{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   error,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, ok := error.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.Response{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	exception, ok := error.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}
