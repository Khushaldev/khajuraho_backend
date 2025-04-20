package services

import (
	"khajuraho/backend/repository"
	"log"
)

func SendOTP() error {
	otp, err := repository.GetOTP()
	if err != nil {
		return err
	}

	log.Println(otp)
	return nil
}
