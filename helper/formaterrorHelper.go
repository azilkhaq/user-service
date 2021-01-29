package helper

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "email_address") {
		return errors.New("Email Already Taken")
	}
	if strings.Contains(err, "phone_number") {
		return errors.New("Phone Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	return errors.New("Incorrect Details")
}
