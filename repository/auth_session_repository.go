package repository

import (
	db "khajuraho/backend/database"
	"khajuraho/backend/model"

	"gorm.io/gorm"
)

func CreateAuthSession(authSession model.AuthSession) (*model.AuthSession, error) {
	result := db.Instance.Create(&authSession)
	if result.Error != nil {
		return nil, result.Error
	}

	return &authSession, nil
}

func FindAuthSessionById(userId string) (*model.AuthSession, error) {
	var AuthSession model.AuthSession
	result := db.Instance.Where("provider_user_id = ?", userId).First(&AuthSession)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return &AuthSession, nil
}
