package services

import (
	"mail"
)

func EmailIsValid (m string) (msg string){
    _, err := mail.ParseAddress(m)
	if err == nil {
		msg = "有効なメールアドレスです"
	} else {
		msg = "不正なメールアドレスです"
	}
	return
}