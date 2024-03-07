package user_usecase

import (
	"EmployeeManagementApp/app/serverconfig"
	"EmployeeManagementApp/domain"
	"EmployeeManagementApp/util"
	"context"
	"time"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUsersUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (u userUsecase) Store(_ context.Context, users *domain.User) (err util.ErrorModel) {
	var (
		fileName = "user_usecase.go"
		funcName = "Store"
		db       = serverconfig.ServerAttribute.DBConnection
		timeNow  = time.Now()
	)

	existedUser, _ := u.userRepo.GetByUsername(nil, db, users.Username)
	if existedUser != (domain.User{}) {
		err = util.GenerateConflictError(fileName, funcName)
		return
	}

	tx, errs := db.Begin()
	if errs != nil {
		return
	}

	defer func() {
		if err.Err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	users.Password, errs = util.HashPassword(users.Password)
	if errs != nil {
		err = util.GenerateInternalDBServerError(fileName, funcName, errs)
		return
	}

	users.CreatedAt = timeNow
	users.UpdatedAt = timeNow

	err = u.userRepo.Store(nil, tx, users)
	if err.Err != nil {
		return
	}

	err = util.GenerateNonError()
	return
}

func (u userUsecase) GetLogin(_ context.Context, users *domain.User) (output domain.User, err util.ErrorModel) {
	db := serverconfig.ServerAttribute.DBConnection
	output, err = u.userRepo.GetLogin(nil, db, users)
	if err.Err != nil {
		return
	}

	return
}
