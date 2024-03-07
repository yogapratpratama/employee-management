package domain

import (
	"EmployeeManagementApp/repository/employee_repository/model"
	param "EmployeeManagementApp/repository/util"
	"EmployeeManagementApp/util"
	"context"
	"database/sql"
	"time"
)

type EmployeeRequest struct {
	GetListParameter
	ID           int64  `json:"id"`
	Name         string `json:"name" binding:"required,min=1,max=255"`
	Nip          string `json:"nip" binding:"required,min=1,max=255"`
	Birthplace   string `json:"birthplace"`
	BirthdateStr string `json:"birthdate" binding:"required"`
	Age          int32  `json:"age"`
	Address      string `json:"address"`
	Religion     string `json:"religion"`
	Gender       string `json:"gender" binding:"required,oneof=laki-laki perempuan"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email" binding:"required,email"`
	CreatedAtStr string `json:"created_at"`
	UpdatedAtStr string `json:"updated_at"`
	DeletedAtStr string `json:"deleted_at"`
	Birthdate    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type EmployeeUpdateRequest struct {
	GetListParameter
	ID           int64  `json:"id"`
	Name         string `json:"name" binding:"required,min=1,max=255"`
	Nip          string `json:"nip" binding:"required,min=1,max=255"`
	Birthplace   string `json:"birthplace"`
	BirthdateStr string `json:"birthdate" binding:"required"`
	Age          int32  `json:"age"`
	Address      string `json:"address"`
	Religion     string `json:"religion"`
	Gender       string `json:"gender" binding:"required,oneof=laki-laki perempuan"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email" binding:"required,email"`
}

type ListEmployeeResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Nip         string `json:"nip"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type ViewEmployeeResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Nip         string    `json:"nip"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Birthplace  string    `json:"birthplace"`
	Birthdate   time.Time `json:"birthdate"`
	Age         int32     `json:"age"`
	Address     string    `json:"address"`
	Religion    string    `json:"religion"`
	Gender      string    `json:"gender"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type EmployeeUsecase interface {
	Add(ctx context.Context, context ContextModel, employee *EmployeeRequest) util.ErrorModel
	Update(ctx context.Context, context ContextModel, employee *EmployeeUpdateRequest) util.ErrorModel
	Delete(ctx context.Context, context ContextModel, employee *EmployeeRequest) util.ErrorModel
	Fetch(ctx context.Context, employee *EmployeeRequest) ([]ListEmployeeResponse, util.ErrorModel)
	GetByID(ctx context.Context, id int64) (ViewEmployeeResponse, util.ErrorModel)
}

type EmployeeRepository interface {
	Add(ctx context.Context, tx *sql.Tx, employee *model.EmployeeModel) util.ErrorModel
	Update(ctx context.Context, tx *sql.Tx, employee model.EmployeeModel) util.ErrorModel
	Delete(ctx context.Context, tx *sql.Tx, employee model.EmployeeModel) util.ErrorModel
	Fetch(ctx context.Context, db *sql.DB, param param.GetListParameterModel) ([]model.EmployeeModel, util.ErrorModel)
	GetByID(ctx context.Context, db *sql.DB, id int64) (model.EmployeeModel, util.ErrorModel)
}
