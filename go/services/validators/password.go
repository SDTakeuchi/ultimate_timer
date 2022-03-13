package services

import (
	"fmt"
)

func PasswordIsValid(p string) (msg string){
	minLen := 6
	maxlen := 30
	ng_password := []byte(
		"password",
		"12345678",
		"qwerty",
	)

	if p == "" {
		msg = "パスワードが入力されていません"
		return
	}
	if len(p) < minLen {
		msg = fmt.Sprintf("文字数が不足しています; %v文字以上入力してください", minLen)
		return
	}
	if len(p) > maxlen {
		msg = fmt.Sprintf("文字数が超過しています; %v文字以下で入力してください", maxLen)
		return
	}
	if Contains(p, ng_password) {
		msg = "単純すぎます"
		return
	}
	return
}