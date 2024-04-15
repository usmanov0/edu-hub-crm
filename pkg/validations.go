package pkg

import (
	"edu-sphere-crm/entity"
	errors2 "edu-sphere-crm/pkg/customErrors"
	"errors"
	"regexp"
	"strings"
)

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (f UserFactory) CreateNewUser(user *entity.NewUser) *entity.User {
	return &entity.User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}
}

func ValidateUserInfoForSignUp(firstName, lastName, email, password, phoneNumber, address string) error {
	if strings.TrimSpace(firstName) == "" {
		return errors2.ErrEmptyFirstName
	}
	if strings.TrimSpace(lastName) == "" {
		return errors2.ErrEmptyLastName
	}
	if errors.Is(ValidateEmail(email), errors2.ErrInvalidEmailFormat) {
		return errors2.ErrEmptyMail
	}
	if errors.Is(validatePassword(password), errors2.ErrInvalidPassword) {
		return errors2.ErrInvalidPassword
	}
	if !IsValidPhoneNumber(phoneNumber) {
		return errors2.ErrPhoneNumberIsInvalid
	}
	return nil
}

func ValidateUserInfoForSignIn(email, password string) error {
	if strings.TrimSpace(email) == "" {
		return errors2.ErrBadCredentials
	}
	if errors.Is(validatePassword(password), errors2.ErrInvalidPassword) {
		return errors2.ErrBadCredentials
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors2.ErrInvalidPassword
	}

	var (
		upperCase = regexp.MustCompile(`[A-Z]`)
		lowerCase = regexp.MustCompile(`[a-z]`)
		digit     = regexp.MustCompile(`[0-9]`)
	)
	if !upperCase.MatchString(password) || !lowerCase.MatchString(password) {
		return errors2.ErrInvalidPassword
	}
	if !digit.MatchString(password) {
		return errors2.ErrInvalidPassword
	}
	return nil
}

func ValidateEmail(email string) error {
	emailReg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailReg.MatchString(email) {
		return errors2.ErrInvalidEmailFormat
	}
	return nil
}

func IsValidPhoneNumber(phoneNumber string) bool {
	phoneRegex := `^\+?[0-9]{8,15}$`

	pattern := regexp.MustCompile(phoneRegex)

	return pattern.MatchString(phoneNumber)
}
