package account

import (
	"context"
	"dungeons_helper/internal/image"
)

type Account struct {
	accessToken string
	Id          int64       `json:"id"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	Nickname    string      `json:"nickname"`
	Avatar      image.Image `json:"avatar"`
}

type CreateAccountReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type LoginAccountReq struct {
	Email    string `json:"Email"`
	Password string `json:"password"`
}

type LoginAccountRes struct {
	accessToken string
	Id          int64       `json:"id"`
	Email       string      `json:"email"`
	Nickname    string      `json:"nickname"`
	Avatar      image.Image `json:"IdAvatar"`
}

type GoogleAcc struct {
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

type Token struct {
	Token string `json:"token"`
}

type RestoreReq struct {
	Email string `json:"email"`
}

type Repository interface {
	CreateAccount(ctx context.Context, account *CreateAccountReq) error
	GetAccountByEmail(ctx context.Context, email string) (*Account, error)
	UpdatePassword(ctx context.Context, account *Account) error
}

type Service interface {
	CreateAccount(c context.Context, req *CreateAccountReq) error
	Login(c context.Context, req *LoginAccountReq) (*LoginAccountRes, error)
	GoogleAuth(c context.Context, req *GoogleAcc) (*LoginAccountRes, error)
	RestorePassword(c context.Context, email string) error
}
