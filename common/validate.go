package common

import "github.com/go-playground/validator/v10"

type XValidator struct {
	Validator *validator.Validate
}

// This is the Validator instance
// for more information see: https://github.com/go-playground/validator
var Validate = validator.New()

func (v XValidator) Validate(data any) {
	errs := Validate.Struct(data)
	if errs != nil {
		var errors []string
		for _, err := range errs.(validator.ValidationErrors) {
			errors = append(errors, err.Error())
		}
		ParamInvalid.Panic(errors...)
	}
}
