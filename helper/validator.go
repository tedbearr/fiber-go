package helper

import "github.com/go-playground/validator/v10"

var Validator = validator.New()

func Validate(i interface{}) error {
	return Validator.Struct(i)
}
