package repository

import (
	"fmt"
	db "khajuraho/backend/database"
	"khajuraho/backend/dto"
	"khajuraho/backend/model"
	"time"

	"firebase.google.com/go/v4/auth"
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

func FetchOrRegisterUser(
	uid string,
	userRecord *auth.UserRecord,
	req dto.GoogleLoginRequest,
) (*model.User, error) {
	session, err := FindAuthSessionById(uid)
	if err != nil {
		return nil, err
	}

	if session != nil {
		return FindUserById(session.UserID.String())
	}

	newUser := model.User{
		DisplayName:    userRecord.DisplayName,
		Email:          userRecord.Email,
		ProfilePicture: userRecord.PhotoURL,
		IsActive:       true,
		IsVerified:     true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
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
		// IP:             req.IP,        // TODO: get IP from Header
		// UserAgent:      req.UserAgent, // TODO: get UserAgent from Header
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = CreateAuthSession(*authSession)
	if err != nil {
		return nil, err
	}

	return user, nil
}
