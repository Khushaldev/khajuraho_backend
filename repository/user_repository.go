package repository

import (
	"fmt"
	db "khajuraho/backend/database"
	"khajuraho/backend/dto"
	"khajuraho/backend/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindUserById(userId string) (*model.User, error) {
	var user model.User
	result := db.Instance.Where("id = ?", userId).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with ID %s not found", userId)
		}
		return nil, result.Error
	}

	return &user, nil
}

func CreateUser(newUser model.User) (*model.User, error) {
	result := db.Instance.Create(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}

func FetchOrRegisterUser(uid string, req dto.GoogleLoginRequest) (*model.User, error) {
	session, err := FindAuthSessionById(uid)
	if err != nil {
		return nil, err
	}

	if session != nil {
		return FindUserById(session.UserID.String())
	}

	newUser := model.User{
		DisplayName: "Khushal Yadav",            // TODO: Get from firebase
		Email:       "khushalyadav90@gmail.com", // TODO: Get from firebase
		IsActive:    true,
		IsVerified:  true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	user, err := CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	authSession := &model.AuthSession{
		ID:             uuid.New(),
		UserID:         user.ID,
		Provider:       "firebase",
		ProviderUserId: uid,
		LoginType:      "google",
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		IP:             req.IP,
		UserAgent:      req.UserAgent,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	_, err = CreateAuthSession(*authSession)
	if err != nil {
		return nil, err
	}

	return user, nil
}
