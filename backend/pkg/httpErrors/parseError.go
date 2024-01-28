package httpErrors

import (
	errs "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ HttpError = (*ParseError)(nil)

type ParseError struct {
	BaseError
	Inner error
}

func NewParseError(message string, inner error) ParseError {
	return ParseError{
		BaseError{
			http.StatusBadRequest,
			"ParseError",
			message,
		},
		inner,
	}
}

func (pe *ParseError) JSON() gin.H {
	return gin.H{
		"Type":   pe.Type,
		"Detail": pe.Message,
		"Inner":  pe.Inner.Error(),
	}
}

func (pe *ParseError) ToError() gin.Error {
	return gin.Error{
		Err:  errs.New(fmt.Sprintf("ParseError: %s", pe.Message)),
		Type: gin.ErrorTypePrivate,
		Meta: pe,
	}
}
