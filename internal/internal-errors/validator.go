package internalerrors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	field := strings.ToLower(validationError.StructField())

	fmt.Println(validationError.Field())
	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "max":
		return errors.New(field + " is required with max " + validationError.Param())
	case "min":
		return errors.New(field + " is required with min " + validationError.Param())
	}

	return nil
}
