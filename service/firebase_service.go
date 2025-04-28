package service

import (
	"context"
	"khajuraho/backend/config"

	"firebase.google.com/go/v4/auth"
)

func VerifyFirebaseUserIDToken(idToken string) (*auth.Token, error) {
	return config.FirebaseAuth.VerifyIDToken(context.Background(), idToken)
}

func GetFirebaseUserDetailsByUID(uid string) (*auth.UserRecord, error) {
	return config.FirebaseAuth.GetUser(context.Background(), uid)

}
