package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
)

type ErrResponse struct {
	Errors []Errors `json:"errors"`
}
type Errors struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func NewValidator() *validator.Validate {
	validate := validator.New()
	// Using the names which have been specified for JSON representations of structs, rather than normal Go field names
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return validate
}
func ToErrResponse(err error) *ErrResponse {

	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		resp := ErrResponse{
			Errors: make([]Errors, len(fieldErrors)),
		}
		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = Errors{
					Field: err.Field(),
					Msg:   fmt.Sprintf("%s is a required field", err.Field()),
				}
			case "min":
				resp.Errors[i] = Errors{
					Field: err.Field(),
					Msg:   fmt.Sprintf("must be a minum of %s in length", err.Param()),
				}
			case "url":
				resp.Errors[i] = Errors{
					Field: err.Field(),
					Msg:   fmt.Sprintf("%s must be a valid URL", err.Field()),
				}
			case "email":
				resp.Errors[i] = Errors{
					Field: err.Field(),
					Msg:   fmt.Sprintf("%s must be a valid email", err.Field()),
				}
			case "eqfield":
				resp.Errors[i] = Errors{
					Field: err.Field(),
					Msg:   fmt.Sprintf("%s confirm password must be equal to password", err.Field()),
				}
			default:
				resp.Errors[i] = Errors{
					Field: err.Field(),
					Msg:   fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Tag()),
				}
			}
		}
		return &resp
	}
	return nil
}
