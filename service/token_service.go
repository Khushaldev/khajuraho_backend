package service

import (
	"khajuraho/backend/config"
	"khajuraho/backend/dto"
	"khajuraho/backend/model"
	"khajuraho/backend/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateJWT(subject string, expiry time.Time) (*model.JWTToken, error) {
	jwtSecret := []byte(config.AppConfig.JWTSecret)

	claims := model.Claims{
		UserID: subject,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	return &model.JWTToken{Token: tokenString, ExpiresAt: claims.ExpiresAt.Time}, nil
}
func GenerateAuthTokens(userID string) (*model.AuthTokens, error) {
	accessToken, err := CreateJWT(userID, time.Now().Add(15*time.Minute))
	if err != nil {
		return nil, err
	}

	refreshTokenID := uuid.NewString()
	refreshToken, err := CreateJWT(refreshTokenID, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return nil, err
	}

	return &model.AuthTokens{
		AccessToken:  model.JWTToken{Token: accessToken.Token, ExpiresAt: accessToken.ExpiresAt},
		RefreshToken: model.JWTToken{Token: refreshToken.Token, ExpiresAt: refreshToken.ExpiresAt},
	}, nil
}

func CreateRefreshSession(userID uuid.UUID, req dto.GoogleLoginRequest, refresh model.JWTToken) error {
	session := model.RefreshSession{
		ID:        uuid.New(),
		UserID:    userID,
		Token:     refresh.Token,
		ExpiresAt: refresh.ExpiresAt,
		Device:    req.Device,
		IP:        req.IP,
		UserAgent: req.UserAgent,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Revoked:   false,
	}
	return repository.SaveRefreshToken(&session)
}

func RefreshTokenService(refreshToken string) (*model.AuthTokens, error) {
	session, err := repository.FindValidRefreshSession(refreshToken)
	if err != nil || session.Revoked || session.ExpiresAt.Before(time.Now()) {
		return nil, err
	}

	tokens, err := GenerateAuthTokens(session.UserID.String())
	if err != nil {
		return nil, err
	}

	err = CreateRefreshSession(session.UserID, dto.GoogleLoginRequest{
		Device:    session.Device,
		IP:        session.IP,
		UserAgent: session.UserAgent,
	}, tokens.RefreshToken)
	if err != nil {
		return nil, err
	}

	err = repository.RevokeRefreshToken(session.ID.String())
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
