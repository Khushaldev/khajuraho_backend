package repository

import (
	"errors"
	db "khajuraho/backend/database"
	"khajuraho/backend/model"
)

// SaveRefreshToken saves the refresh session in the DB.
func SaveRefreshToken(session *model.RefreshSession) error {
	return db.Instance.Create(session).Error
}

// FindValidRefreshSession returns a valid refresh session if not revoked and hasn't expired.
func FindValidRefreshSession(refreshToken string) (*model.RefreshSession, error) {
	var session model.RefreshSession
	err := db.Instance.
		Where("token = ? AND revoked = false AND expires_at > NOW()", refreshToken).
		First(&session).Error

	if err != nil {
		return nil, err
	}
	return &session, nil
}

// RevokeRefreshToken marks the refresh token as revoked.
func RevokeRefreshToken(refreshToken string) error {
	result := db.Instance.Model(&model.RefreshSession{}).
		Where("token = ?", refreshToken).
		Update("revoked", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no session found to revoke")
	}

	return nil
}
