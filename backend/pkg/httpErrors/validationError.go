package httpErrors

import (
	errs "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var _ HttpError = (*ValidationError)(nil)

type ValidationError struct {
	BaseError
	Errors map[string][]string
}

func NewValidationError() *ValidationError {
	ve := ValidationError{}
	ve.SetDefaults()
	return &ve
}

func (ve *ValidationError) SetDefaults() {
	ve.HTTPStatus = http.StatusBadRequest
	ve.Type = "ValidationError"
	ve.Message = "The input provided is invalid."
}

func (ve *ValidationError) addField(field, element string) {
	if ve.Errors == nil {
		ve.Errors = make(map[string][]string)
	}
	if ve.Errors[field] == nil {
		ve.Errors[field] = make([]string, 0)
	}
	ve.Errors[field] = append(
		ve.Errors[field],
		element,
	)
}

func (ve *ValidationError) AddRequiredField(field string) {
	element := fmt.Sprintf("%s is an required field", field)
	ve.addField(field, element)
}

func (ve *ValidationError) AddInvalidField(field string) {
	element := fmt.Sprintf("%s is invalid", field)
	ve.addField(field, element)
}

func (ve *ValidationError) AddInvalidFieldWithContext(field, context string) {
	element := fmt.Sprintf("%s is invalid- %s", field, context)
	ve.addField(field, element)
}

// CheckStringValidationWithMinimumLength Check if the string meets the minimum length
// Adds a required field error if the field is empty
// Adds an invalid field error if the field does not meet the minimum required length
func (ve *ValidationError) CheckStringValidationWithMinimumLength(
	field string, value string, minLength int,
) {
	if len(value) == 0 {
		ve.AddRequiredField(field)
	} else if len(value) <= minLength {
		ve.AddInvalidFieldWithContext(
			field,
			fmt.Sprintf(
				"%s must contain at least 2 characters",
				strings.ToTitle(field),
			),
		)
	}
}

func (ve *ValidationError) IsEmpty() bool {
	return len(ve.Errors) == 0
}

func (ve *ValidationError) JSON() gin.H {
	return gin.H{
		"Type":   ve.Type,
		"Detail": ve.Errors,
	}
}

func (ve *ValidationError) ToError() gin.Error {
	if len(ve.Errors) == 0 {
		return gin.Error{}
	}
	ve.SetDefaults()
	return gin.Error{
		Err:  errs.New(fmt.Sprintf("ValidationError: %s", ve.Message)),
		Type: gin.ErrorTypePrivate,
		Meta: ve,
	}
}
