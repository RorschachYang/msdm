package dao

import (
	"errors"

	"gorm.io/gorm"
)

func CreateUser(user *User) error {
	return db.Create(user).Error
}

func UpdatetUser(user *User) error {
	var existingUser User
	err := db.Where("open_id = ?", user.OpenID).First(&existingUser).Error
	if err != nil {
		return err
	}

	existingUser.AvatarURL = user.AvatarURL
	existingUser.NickName = user.NickName
	existingUser.OpenID = user.OpenID

	return db.Save(&existingUser).Error
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
