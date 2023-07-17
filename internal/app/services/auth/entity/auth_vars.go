package entity

import (
	"strings"

	"github.com/nhatth/api-service/internal/app/helpers"
)

const TYPE_TAG_JSON = "json"

type AuthEmailPassword struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type AuthRegister struct {
	FirstName string `json:"firstName" form:"firstName"`
	LastName  string `json:"lastName" form:"lastName"`
	Salt      string
	AuthEmailPassword
}

func (authen *AuthEmailPassword) Validate() (map[string]string, bool) {

	authen.Email = strings.TrimSpace(authen.Email)

	var errors = make(map[string]string)

	errorValidate := false

	if !emailIsValid(authen.Email) {

		errorValidate = true

		jsonTag := helpers.FindStructFieldJSONName(authen, &authen.Email, TYPE_TAG_JSON)

		errors[jsonTag] = ErrEmailIsNotValid.Error()
	}

	authen.Password = strings.TrimSpace(authen.Password)

	if err := checkPassword(authen.Password); err != nil {

		errorValidate = true

		jsonTag := helpers.FindStructFieldJSONName(authen, &authen.Password, TYPE_TAG_JSON)

		errors[jsonTag] = err.Error()
	}

	return errors, errorValidate
}

func (authRegis *AuthRegister) Validate() (map[string]string, bool) {

	errors, errorValidate := authRegis.AuthEmailPassword.Validate()

	authRegis.FirstName = strings.TrimSpace(authRegis.FirstName)

	if err := checkFirstName(authRegis.FirstName); err != nil {

		jsonTag := helpers.FindStructFieldJSONName(authRegis, &authRegis.FirstName, TYPE_TAG_JSON)

		errors[jsonTag] = err.Error()
	}

	authRegis.LastName = strings.TrimSpace(authRegis.LastName)

	if err := checkLastName(authRegis.LastName); err != nil {

		jsonTag := helpers.FindStructFieldJSONName(authRegis, &authRegis.LastName, TYPE_TAG_JSON)

		errors[jsonTag] = err.Error()
	}

	return errors, errorValidate
}
