package passwd

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var cost = bcrypt.DefaultCost
var passwordMaxLength = 128

func Encode(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func EncodeString(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(pwd), nil
}

func Match(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func MatchString(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Filter(password string) (string, error) {
	if len(password) > passwordMaxLength {
		return "", errors.New("密码太长")
	}
	return password, nil
}
