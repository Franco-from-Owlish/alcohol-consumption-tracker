package httpErrors

import "github.com/gin-gonic/gin"

type BaseError struct {
	HTTPStatus int
	Type       string
	Message    string
}

type HttpError interface {
	JSON() gin.H
	ToError() gin.Error
}
