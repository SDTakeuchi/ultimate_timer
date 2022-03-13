package models

import (
	"ultimate_timer/services"
	"errors"
	"time"
)

type User struct {
	BaseModel
	Email       string `db:"email" json:"email"`
	Password    string `db:"memo" json:"memo"`
	UserSetting UserSetting
}

type UserSetting struct {
	SoundOn bool `db:"sound_on" json:"sound_on"`
	UserID  string
}

/*
constructor is for preparing objects (returns struct and error)
setter is for updating new objects (returns error if failed to udpate)
*/

// constructor
func NewUser(email, password string) (*User, error) {
	msg := services.EmailIsValid(email)
	if msg != nil {
		return nil, errors.New(msg)
	}
	msg = services.PasswordIsValid(password)
	if msg != nil {
		return nil, errors.New(msg)
	}

	if name == "" {
		return nil, errors.New("項目名を入力してください")
	}

	now := time.Now()
	id := lib.GenUuid()

	item := &User{
		BaseModel: BaseModel{
			ID:        id,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Email:         email,
		Memo:         memo,
		Price:        price,
		PurchaseDate: purchaseDate,
		// SmallCategoryId: smallCategoryId,
		// UserId:          userId,
	}

	return item, nil
}

// setter
func (i *User) Set(
	name, memo string,
	price int,
	purchaseDate time.Time) error {

	if name == "" {
		return errors.New("項目名を入力してください")
	}

	now := time.Now()

	i.UpdatedAt = now
	i.Name = name
	i.Memo = memo
	i.Price = price
	i.PurchaseDate = purchaseDate
	// i.SmallCategoryId = smallCategoryId

	return nil
}
