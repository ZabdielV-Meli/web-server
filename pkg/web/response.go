package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessfulResponse struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Success escribe una respuesta exitosa
func Success(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, SuccessfulResponse{
		Data: data,
	})
}

// Failure escribe una respuesta fallida
func Failure(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, ErrorResponse{
		Message: err.Error(),
		Status:  status,
		Code:    http.StatusText(status),
	})
}
