package account

import (
	"context"
	"dungeons_helper/util"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateAccount(c context.Context, req *CreateAccountReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return err
	}

	account := &CreateAccountReq{
		Email:    req.Email,
		Password: hashedPassword,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
	}

	err = s.Repository.CreateAccount(ctx, account)
	if err != nil {
		return err
	}

	err = sendWelcomeEmail(req.Email)
	if err != nil {
		return err
	}

	return nil
}

func sendEmail(toEmail string, subject string, body string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	headers := make(map[string]string)
	headers["From"] = smtpUsername
	headers["To"] = toEmail
	headers["Subject"] = subject

	message := ""
	for key, value := range headers {
		message += key + ": " + value + "\r\n"
	}
	message += "\r\n" + body

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	serverAddr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	err := smtp.SendMail(serverAddr, auth, smtpUsername, []string{toEmail}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

func sendWelcomeEmail(toEmail string) error {
	subject := "Добро пожаловать!"
	body := "Спасибо за регистрацию!"

	err := sendEmail(toEmail, subject, body)
	if err != nil {
		return err
	}

	return nil
}
