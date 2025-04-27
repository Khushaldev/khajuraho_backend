package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App
var FirebaseAuth *auth.Client

func InitFirebase() {
	opt := option.WithCredentialsFile("./khajuraho-private-key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("failed to initialize firebase app: %v", err)
	}
	FirebaseApp = app

	FirebaseAuth, err = app.Auth(context.Background())

	if err != nil {
		log.Fatalf("failed to initialize firebase auth: %v", err)
	}
}
