package httpErrors

import (
	errs "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ HttpError = (*DatabaseError)(nil)

type DatabaseError struct {
	BaseError
	Inner error
}

func NewDatabaseError(cause string, inner error) DatabaseError {
	return DatabaseError{
		BaseError{
			http.StatusInternalServerError,
			"DatabaseError",
			cause,
		},
		inner,
	}
}

func (dbe *DatabaseError) JSON() gin.H {
	return gin.H{
		"Type":   dbe.Type,
		"Detail": dbe.Message,
		"Inner":  dbe.Inner.Error(),
	}
}

func (dbe *DatabaseError) ToError() gin.Error {
	return gin.Error{
		Err:  errs.New(fmt.Sprintf("DatabaseError: %s", dbe.Message)),
		Type: gin.ErrorTypePrivate,
		Meta: dbe,
	}
}
