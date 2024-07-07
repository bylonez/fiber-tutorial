package valid

import (
	"github.com/bylonez/fiber-tutorial/pkg/ex"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3/log"
)

var validate = validator.New()

func Valid(data any) {
	errs := validate.Struct(data)
	if errs != nil {
		var errors []any
		for _, err := range errs.(validator.ValidationErrors) {
			errors = append(errors, err.Error())
		}
		ex.ParamInvalid.Panic(errors...)
	}
}

func Parse(data any, f func(out interface{}) error) {
	err := f(data)
	if err != nil {
		log.Warn(err)
		ex.ParamInvalid.Panic()
	}
	Valid(data)
}

func Register(tag string, fn validator.Func, callValidationEvenIfNull ...bool) {
	err := validate.RegisterValidation(tag, fn, callValidationEvenIfNull...)
	if err != nil {
		log.Fatal(err)
	}
}
