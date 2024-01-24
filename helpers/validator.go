package helpers

import (
	"fmt"
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	trans    ut.Translator
	validate *validator.Validate
)

func InitValidator() {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ = uni.GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
}

func ValidateVariable(f interface{}, tag string, errorMessage string) error {
	if err := validate.Var(f, tag); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			return fmt.Errorf(e.Translate((trans)))
		}
	}

	return nil
}

func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			return fmt.Errorf(e.Translate(trans))
		}
	}

	return nil
}

func IsValidPhoneNumber(phoneNumber string) bool {
	pattern := `^\+7 \(\d{3}\) \d{3}-\d{2}-\d{2}-\d{2}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(phoneNumber)
}
