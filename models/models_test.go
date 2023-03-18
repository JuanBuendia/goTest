package models

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/go-playground/locales/es"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"

	es_translations "github.com/go-playground/validator/v10/translations/es"
)

type Site2 struct {
	Id          uint
	CategoryID  int     `json:"category_id"  binding:"required"`
	Name        string  `json:"name"   validate:"required"`
	Description string  `json:"description"   validate:"required"`
	Lat         float32 `json:"lat" sql:"type:decimal(10,5);"`
	Lon         float32 `json:"lon" sql:"type:decimal(10,5);"`
	OfficeHours string  `json:"office_hours" `
	Phone       string  `json:"phone" `
	WebSite     string  `json:"website" `
	Thumbnail   string  `json:"thumbnail"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (site *Site2) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(site)
}

func TestValidation(t *testing.T) {
	site := &Site2{
		Name:        "nombre",
		Description: "des",
		SKU:         "abs-abc-def",
	}

	err := site.Validate()
	if err != nil {
		t.Fatal(err)
	}

}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-absd-dfsdf
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true

}

type person struct {
	Name                string `validate:"required,min=4,max=15"`
	Email               string `validate:"required,email"`
	Age                 int    `validate:"required,numeric,min=18"`
	DriverLicenseNumber string `validate:"omitempty,len=12,numeric"`
}

func TestValidationPersonOK(t *testing.T) {

	p := person{
		Name:                "Joes",
		Email:               "dummyemailgmail.com",
		Age:                 19,
		DriverLicenseNumber: "",
	}
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		t.Fatal(err)
	}

}

type Users struct {
	Name   string `form:"name" json:"name" validate:"required,CustomValidationErrors"` // Contains custom functions
	Age    uint8  `form:"age" json:"age" validate:"required,gt=18"`
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" validate:"required,len=6"`
}

type Users2 struct {
	Name   string `form:"name" json:"name" validate:"required"`
	Age    uint8  `form:"age" json:"age" validate:"required,gt=18"`
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" validate:"required,len=6"`
}

func TestUser(t *testing.T) {

	users := &Users{
		Name:   "admin",
		Age:    12,
		Passwd: "123",
		Code:   "123456",
	}
	validate := validator.New()
	// Register custom functions
	_ = validate.RegisterValidation("CustomValidationErrors", CustomValidationErrors)
	err := validate.Struct(users)
	if err != nil {
		t.Fatal(err)
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err) //Key: 'Users.Name' Error:Field validation for 'Name' failed on the 'CustomValidationErrors' tag
			return
		}
	}

}

func CustomValidationErrors(fl validator.FieldLevel) bool {
	return fl.Field().String() != "admin"
}

func TestTranslation(t *testing.T) {

	fmt.Println("___________________________")
	users := &Users2{
		Name:   "admin",
		Age:    12,
		Passwd: "123",
		Code:   "126",
	}
	uni := ut.New(es.New())
	trans, _ := uni.GetTranslator("es")
	validate := validator.New()
	// Validators register translators
	err := es_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(users)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Translate(trans)) //Age Must be greater than 18

		}
	}

}
