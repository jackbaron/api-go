package entity

import "strings"

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

func (authen *AuthEmailPassword) Validate() error {
	authen.Email = strings.TrimSpace(authen.Email)

	if !emailIsValid(authen.Email) {

		return ErrEmailIsNotValid
	}

	authen.Password = strings.TrimSpace(authen.Password)

	if err := checkPassword(authen.Password); err != nil {

		return err
	}

	return nil
}

func (authRegis *AuthRegister) Validate() error {

	if err := authRegis.AuthEmailPassword.Validate(); err != nil {

		return err
	}

	authRegis.FirstName = strings.TrimSpace(authRegis.FirstName)

	if err := checkFirstName(authRegis.FirstName); err != nil {

		return err
	}

	authRegis.LastName = strings.TrimSpace(authRegis.LastName)

	if err := checkLastName(authRegis.LastName); err != nil {

		return err
	}

	return nil
}
