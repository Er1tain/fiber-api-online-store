package validation

import "net/mail"

func CheckEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
