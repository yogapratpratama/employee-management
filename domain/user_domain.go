package domain

import (
	"EmployeeManagementApp/util"
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username" binding:"required,min=4,max=20"`
	Password     string `json:"password" binding:"required,min=8,max=20"`
	CreatedAtStr string `json:"created_at"`
	UpdatedAtStr string `json:"updated_at"`
	DeletedAtStr string `json:"deleted_at"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	Deleted      bool `json:"deleted"`
}

type Login struct {
	Username string `json:"username" binding:"required,min=4,max=20"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type UserUsecase interface {
	Store(ctx context.Context, users *User) util.ErrorModel
	GetLogin(ctx context.Context, users *User) (User, util.ErrorModel)
}

type UserRepository interface {
	Store(ctx context.Context, tx *sql.Tx, users *User) util.ErrorModel
	GetByUsername(ctx context.Context, db *sql.DB, username string) (User, util.ErrorModel)
	GetLogin(ctx context.Context, db *sql.DB, users *User) (User, util.ErrorModel)
}
