package validators

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func Descriptive(verr validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: f.Field(), Reason: err})

	}

	return errs
}

func ValidateInput(c *gin.Context, input interface{}) error {
	unError := errors.New("error en la validacion")
	validate := validator.New()
	if err := c.ShouldBind(input); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": Descriptive(verr)})
			return unError
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return unError
	}
	err := validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return unError
	}
	return nil
}
