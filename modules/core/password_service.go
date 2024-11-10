package core

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrPasswordVerify = errors.New("password are not equal")

type PasswordService interface {
	Generate(password string) (string, error)
	Verify(password string, hash string) error
}

type PasswordServiceImpl struct{}

func (p PasswordServiceImpl) Generate(password string) (string, error) {
	var (
		b   []byte
		err error
	)

	if b, err = bcrypt.GenerateFromPassword([]byte(password), 14); err != nil {
		return string(b), err
	}

	return string(b), nil
}

func (p PasswordServiceImpl) Verify(password string, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return ErrPasswordVerify
	}

	return nil
}

func NewPasswordService() PasswordService {
	return &PasswordServiceImpl{}
}
