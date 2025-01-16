package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

type Errors struct {
	Field   string `json:"field" exemple:"name"`
	Message string `json:"message" exemple:"name is required"`
}

var date validator.Func = func(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)

	if ok {
		layout := "2006-01-02"
		_, err := time.Parse(layout, value)

		if err != nil {
			return false
		}
	}
	return true
}

var dateTime validator.Func = func(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)

	if ok {
		layoutWithTimeZone := "2006-01-02T15:04:05Z"
		_, err := time.Parse(layoutWithTimeZone, value)

		if err == nil {
			return true
		}

		layoutWithOutTimeZone := "2006-01-02T15:04:05"
		_, err = time.Parse(layoutWithOutTimeZone, value)

		if err == nil {
			return true
		}
	}
	return true
}

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		value.RegisterValidation("cdate", date)
		value.RegisterValidation("cdatetime", dateTime)
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(value, transl)
	}
}

func ValidatorError(validateError error) []Errors {

	var jsonValidatorError validator.ValidationErrors
	var jsonError *json.UnmarshalFieldError

	errorsCauses := []Errors{}
	if errors.As(validateError, &jsonError) {
		e := validateError.(*json.UnmarshalTypeError)
		cause := Errors{
			Field:   e.Field,
			Message: fmt.Sprintf("it can not be %s", e.Value),
		}

		errorsCauses = append(errorsCauses, cause)

	} else if errors.As(validateError, &jsonValidatorError) {
		for _, e := range validateError.(validator.ValidationErrors) {
			cause := Errors{
				Field:   e.Translate(transl),
				Message: e.Namespace(),
			}

			errorsCauses = append(errorsCauses, cause)
		}
	} else {
		cause := Errors{
			Field:   fmt.Sprintf("Please check your payload %s", validateError.Error()),
			Message: "Payload",
		}

		errorsCauses = append(errorsCauses, cause)
	}

	return errorsCauses
}
