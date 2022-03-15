package services

import (
	"mail"
)

func EmailIsValid (m string) (errMsg string){
    _, err := mail.ParseAddress(m)
	if err != nil {
		errMsg = "不正なメールアドレスです"
	}
	return
}