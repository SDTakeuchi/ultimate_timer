// package models

// import (
// 	"time"
// 	"ultimate_timer/services"
// )

// type UserSetting struct {
// 	BaseModel
// 	SoundOn bool `db:"sound_on" json:"sound_on"`
// 	UserID  string
// }

// // constructor
// func NewUserSetting(userID string) (*UserSetting, error) {
// 	now := time.Now()
// 	id := services.GenUuid()

// 	us := &UserSetting{
// 		BaseModel: BaseModel{
// 			ID:        id,
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 		SoundOn:         true,
// 		UserID: userID,
// 	}

// 	return us, nil
// }

// // setter
// func (us *UserSetting) Set(soundOn bool) error {
// 	now := time.Now()
// 	us.UpdatedAt = now
// 	us.SoundOn = soundOn
// 	return nil
// }
