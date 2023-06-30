package utils

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var (
	validate *validator.Validate
)

type HandleErrror struct {
	Tag     string
	Message string
}

func PassswordHash(password string) []byte {
	hashing, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return hashing
}

func ComparePassword(pass_has string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pass_has), []byte(pass))

	return err == nil
}

func Validate(data interface{}) []HandleErrror {
	var errors []HandleErrror
	validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationErr := range validationErrors {
			errors = append(errors, HandleErrror{Tag: validationErr.Field(), Message: validationErr.Tag()})
		}
	}
	return errors
}
