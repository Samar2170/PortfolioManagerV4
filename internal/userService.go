package internal

import (
	"context"

	"github.com/samar2170/portfolio-manager-v4/client/cognitio/cauth"
	"github.com/samar2170/portfolio-manager-v4/internal/models"
	"github.com/samar2170/portfolio-manager-v4/internal/utils"
)

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
	if err != nil {
		return err
	}
	err = createGeneralAccountForUser(&dbUser)
	return err
}

func createGeneralAccountForUser(user *models.User) error {
	generalAccount := models.GeneralAccount{
		UserCID:     user.UserCID,
		AccountCode: user.UserCID,
	}
	return models.CreateModelInstance(&generalAccount)
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

func RegisterBankAccount(ba BankAccountRequest, userCID string) error {
	_, err := models.GetUserByCID(userCID)
	if err != nil {
		return err
	}
	bankAccount := models.BankAccount{
		UserCID:   userCID,
		Bank:      ba.Bank,
		AccountNo: ba.AccountNo,
	}
	return models.CreateModelInstance(&bankAccount)

}

func RegisterDematAccount(da DematAccountRequest, userCID string) error {
	_, err := models.GetUserByCID(userCID)
	if err != nil {
		return err
	}
	dematAccount := models.DematAccount{
		UserCID:     userCID,
		AccountCode: da.AccountCode,
		Broker:      da.Broker,
	}
	return models.CreateModelInstance(&dematAccount)
}
