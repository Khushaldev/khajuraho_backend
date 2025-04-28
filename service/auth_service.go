package service

import (
	"khajuraho/backend/dto"
	"khajuraho/backend/model"
	"khajuraho/backend/repository"
)

func LoginWithGoogle(req dto.GoogleLoginRequest) (*model.User, *model.AuthTokens, error) {
	firebaseToken, err := VerifyFirebaseUserIDToken(req.IDToken)
	if err != nil {
		return nil, nil, err
	}

	uid := firebaseToken.UID
	userRecord, err := GetFirebaseUserDetailsByUID(uid)
	if err != nil {
		return nil, nil, err
	}

	user, err := repository.FetchOrRegisterUser(uid, userRecord, req)
	if err != nil {
		return nil, nil, err
	}

	tokens, err := GenerateAuthTokens(user.ID.String())
	if err != nil {
		return nil, nil, err
	}

	err = CreateRefreshSession(user.ID, req, tokens)
	if err != nil {
		return nil, nil, err
	}

	return user, tokens, nil
}

func LogoutService(refreshToken string) error {
	return repository.RevokeRefreshToken(refreshToken)
}
