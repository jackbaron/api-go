package entity

import (
	"net/mail"
	"unicode"
)

func emailIsValid(s string) bool {
	_, err := mail.ParseAddress(s)

	return err == nil
}

func checkPassword(s string) error {
	if len(s) < 8 || len(s) > 30 {
		return ErrPasswordIsNotValid
	}

	var (
		hasUper   = false
		hasLower  = false
		hasNumber = false
		hasSecial = false
	)

	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSecial = true
		}
	}

	if !hasUper || !hasLower || !hasSecial || !hasNumber {
		return ErrPasswordIsNotStrong
	}

	return nil
}

func checkFirstName(s string) error {
	if s == "" {
		return ErrFirstNameIsEmpty
	}

	if len(s) > 255 {
		return ErrFirstNameTooLong
	}

	return nil
}

func checkLastName(s string) error {
	if s == "" {
		return ErrLastNameIsEmpty
	}

	if len(s) > 255 {
		return ErrLastNameTooLong
	}

	return nil
}
