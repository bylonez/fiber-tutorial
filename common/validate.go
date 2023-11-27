package common

import (
	"github.com/go-playground/validator/v10"
	"log"
)

var Validate = validator.New()

func Valid(data any) {
	errs := Validate.Struct(data)
	if errs != nil {
		var errors []string
		for _, err := range errs.(validator.ValidationErrors) {
			errors = append(errors, err.Error())
		}
		ParamInvalid.Panic(errors...)
	}
}

func init() {
	// Custom struct validation tag format
	err := Validate.RegisterValidation("teener", func(fl validator.FieldLevel) bool {
		// User.Age needs to fit our needs, 12-18 years old.
		return fl.Field().Int() >= 12 && fl.Field().Int() <= 18
	})
	if err != nil {
		log.Fatal(err)
	}
}
