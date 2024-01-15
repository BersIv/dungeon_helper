package account

import (
	"context"
	"dungeons_helper/util"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
		timeout:    time.Duration(2) * time.Second,
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

	// err = sendWelcomeEmail(req.Email)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (s *service) Login(c context.Context, req *LoginAccountReq) (*LoginAccountRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	account, err := s.Repository.GetAccountByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(req.Password, account.Password)
	if err != nil {
		return nil, err
	}

	ss, err := newToken(account)
	if err != nil {
		return nil, err
	}

	return &LoginAccountRes{accessToken: ss, Id: account.Id, Email: account.Email, Nickname: account.Nickname, Avatar: account.Avatar}, nil
}

func (s *service) GoogleAuth(c context.Context, req *GoogleAcc) (*LoginAccountRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	account, err := s.Repository.GetAccountByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	ss, err := newToken(account)
	if err != nil {
		return nil, err
	}

	return &LoginAccountRes{accessToken: ss, Id: account.Id, Email: account.Email, Nickname: account.Nickname, Avatar: account.Avatar}, nil
}

func (s *service) RestorePassword(c context.Context, req *RestoreReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	account, err := s.Repository.GetAccountByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	hashedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	account.Password = hashedPassword
	if err := s.Repository.UpdatePassword(ctx, account); err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateNickname(c context.Context, req *UpdateNicknameReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	account, err := s.Repository.GetAccountById(ctx, req.Id)
	if err != nil {
		return err
	}

	account.Nickname = req.Nickname
	err = s.Repository.UpdateNickname(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdatePassword(c context.Context, req *UpdatePasswordReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	account, err := s.Repository.GetAccountById(ctx, req.Id)
	if err != nil {
		return err
	}

	err = util.CheckPassword(req.OldPassword, account.Password)
	if err != nil {
		return err
	}

	hashedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	account.Password = hashedPassword
	err = s.Repository.UpdatePassword(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateAvatar(c context.Context, req *UpdateAvatarReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	account, err := s.Repository.GetAccountById(ctx, req.Id)
	if err != nil {
		return err
	}

	account.Avatar = req.Avatar
	err = s.Repository.UpdateAvatar(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

func newToken(res *Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, util.MyJWTClaims{
		Id:       res.Id,
		Nickname: res.Nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(res.Id)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	})
	secretKey := os.Getenv("SECRET_KEY")
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return ss, nil
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
