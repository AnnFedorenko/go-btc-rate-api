package service

import (
	"errors"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

const FileName = "subscribers.txt"

var ErrEmailSubscribed = errors.New("Email already subscribed")
var ErrEmailNotValid = errors.New("Email address is not valid")

type Email struct {
	Address string `form:"email" binding:"required"`
}

func AddEmail(newEmail Email) error {

	if !isEmailValid(newEmail.Address) {
		return ErrEmailNotValid
	}

	isEmailSubscribed, err := isEmailSubscribed(newEmail.Address)
	if err != nil {
		return err
	}
	if isEmailSubscribed {
		return ErrEmailSubscribed
	}

	return appendToFile(newEmail.Address, FileName)
}

func isEmailSubscribed(address string) (bool, error) {
	isFileExist, err := isFileExist(FileName)
	if err != nil {
		return false, err
	}
	if !isFileExist {
		return isFileExist, nil
	}

	isEmailExist, err := isStringExist(address, FileName)
	if err != nil {
		return isEmailExist, errors.New("Fail to check if email already subscribed")
	}
	return isEmailExist, nil
}

func isEmailValid(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}

func SendEmails() error {
	godotenv.Load()
	user := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")
	addr := "smtp.gmail.com:587"
	host := "smtp.gmail.com"
	from := "BTC rate app"

	to := readFileToArray(FileName)
	if to == nil {
		return nil
	}

	rate, err := GetRateFromBinance()
	if err != nil {
		return err
	}

	msg := []byte("From: Bitcoin rate helper\r\n" +
		"Subject: BTCUAH Rate\r\n\r\n" +
		rate.Price +
		"\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	if err = smtp.SendMail(addr, auth, from, to, msg); err != nil {
		return err
	}

	return nil
}
