package dao

import (
	"errors"

	"gorm.io/gorm"
)

func UpsertUser(openid string, nickName string, avatarURL string) error {
	var user User
	err := db.Where("open_id = ?", openid).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == nil {
		user.NickName = nickName
		user.AvatarURL = avatarURL
		return db.Save(&user).Error
	}
	newUser := User{OpenID: openid, NickName: nickName, AvatarURL: avatarURL}
	return db.Create(&newUser).Error
}

func GetUserByOpenID(openID string) (*User, error) {
	var user User
	result := db.Where("open_id = ?", openID).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
