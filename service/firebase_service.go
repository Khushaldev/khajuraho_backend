package service

import (
	"context"
	"khajuraho/backend/config"

	"firebase.google.com/go/v4/auth"
)

func VerifyFirebaseIDToken(IDToken string) (*auth.Token, error) {
	return config.FirebaseAuth.VerifyIDToken(context.Background(), IDToken)
}
