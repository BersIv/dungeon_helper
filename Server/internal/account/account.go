package account

import (
	"context"
)

type Account struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type CreateAccountReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type LoginAccountReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginAccountRes struct {
	accessToken string
	Id          int64  `json:"id"`
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
}

type GoogleAcc struct {
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

type Token struct {
	Token string `json:"token"`
}

type RestoreReq struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

type UpdateNicknameReq struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
}

type UpdatePasswordReq struct {
	Id int64 `json:"id"`
	//	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type UpdateAvatarReq struct {
	Id     int64  `json:"id"`
	Avatar string `json:"avatar"`
}

type Repository interface {
	CreateAccount(ctx context.Context, account *CreateAccountReq) error
	GetAccountByEmail(ctx context.Context, email string) (*Account, error)
	GetAccountById(ctx context.Context, id int64) (*Account, error)
	UpdatePassword(ctx context.Context, account *Account) error
	UpdateNickname(ctx context.Context, account *Account) error
	UpdateAvatar(ctx context.Context, account *Account) error
}

type Service interface {
	CreateAccount(c context.Context, req *CreateAccountReq) error
	Login(c context.Context, req *LoginAccountReq) (*LoginAccountRes, error)
	GoogleAuth(c context.Context, req *GoogleAcc) (*LoginAccountRes, error)
	RestorePassword(c context.Context, req *RestoreReq) error
	UpdateNickname(c context.Context, req *UpdateNicknameReq) error
	UpdatePassword(c context.Context, req *UpdatePasswordReq) error
	UpdateAvatar(c context.Context, req *UpdateAvatarReq) error
}
