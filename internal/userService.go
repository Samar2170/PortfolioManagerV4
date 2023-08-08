package internal

import (
	"context"

	"github.com/samar2170/portfolio-manager-v4/client/cognitio/cauth"
	"github.com/samar2170/portfolio-manager-v4/internal/models"
	"github.com/samar2170/portfolio-manager-v4/internal/utils"
)

type SignupRequest struct {
	Username string
	Password string
	Email    string
}

type LoginRequest struct {
	Username string
	Password string
}

func Signup(s SignupRequest) error {
	encryptedPassword, err := utils.EncryptPassword(s.Password, passwordDecryptionKey)
	if err != nil {
		return err
	}
	resp, err := (*cognitioClient).Signup(context.Background(), &cauth.SignupRequest{
		Username: s.Username,
		Password: encryptedPassword,
		Email:    s.Email,
	})
	if err != nil {
		return err
	}
	dbUser := models.User{
		Username: s.Username,
		UserCID:  resp.Response,
	}
	err = models.CreateModelInstance(&dbUser)
	return err
}

func Login(l LoginRequest) (string, error) {
	encryptedPassword, err := utils.EncryptPassword(l.Password, passwordDecryptionKey)
	if err != nil {
		return "", err
	}
	resp, err := (*cognitioClient).Login(context.Background(), &cauth.LoginRequest{
		Username: l.Username,
		Password: encryptedPassword,
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}
