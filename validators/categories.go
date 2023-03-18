package validators

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateCoolTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "is-cool")
}

func IsISO8601Date(fl validator.FieldLevel) bool {
	regex := "^(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])((?:T|\\s)(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])?(Z)?)?$"
	ISO8601DateRegex := regexp.MustCompile(regex)

	return ISO8601DateRegex.MatchString(fl.Field().String())
}
