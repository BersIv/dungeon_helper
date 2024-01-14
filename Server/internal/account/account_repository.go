package account

import (
	"context"
	"dungeons_helper/db"
)

type repository struct {
	db db.DatabaseTX
}

func NewRepository(db db.DatabaseTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateAccount(ctx context.Context, req *CreateAccountReq) error {
	query := "INSERT INTO image(image) VALUES (?)"
	result, err := r.db.ExecContext(ctx, query, req.Avatar)
	if err != nil {
		return err
	}
	idImage, err := result.LastInsertId()
	if err != nil {
		return err
	}

	query = "INSERT INTO account(email, password, nickname, idAvatar) VALUES (?, ?, ?, ?)"
	_, err = r.db.ExecContext(ctx, query, req.Email, req.Password, req.Nickname, idImage)
	if err != nil {
		return err
	}

	return nil
}
