package account

import (
	"context"
	"dungeons_helper/internal/image"
)

type Account struct {
	Id       int64       `json:"id"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Nickname string      `json:"nickname"`
	Avatar   image.Image `json:"avatar"`
}

type CreateAccountReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type Repository interface {
	CreateAccount(ctx context.Context, account *CreateAccountReq) error
}

type Service interface {
	CreateAccount(c context.Context, req *CreateAccountReq) error
}
