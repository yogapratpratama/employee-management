package employee_usecase

import (
	"EmployeeManagementApp/app/serverconfig"
	"EmployeeManagementApp/domain"
	"EmployeeManagementApp/repository/employee_repository/model"
	"EmployeeManagementApp/util"
	"context"
	"database/sql"
	"time"
)

func (t employeeUsecase) Add(_ context.Context, context domain.ContextModel, employee *domain.EmployeeRequest) (err util.ErrorModel) {
	var (
		fileName      = "employee_usecase_add.go"
		funcName      = "Add"
		db            = serverconfig.ServerAttribute.DBConnection
		timeNow       = time.Now()
		employeeModel model.EmployeeModel
		birthdate     time.Time
	)

	tx, errs := db.Begin()
	if errs != nil {
		err = util.GenerateUnknownServerError(fileName, funcName, errs)
		return
	}

	defer func() {
		if err.Err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	birthdate, errs = time.Parse("2006-01-02T15:04:05Z", employee.BirthdateStr)
	if errs != nil {
		err = util.GenerateUnknownServerError(fileName, funcName, errs)
		return
	}

	employeeModel = model.EmployeeModel{
		Name:        sql.NullString{String: employee.Name},
		Nip:         sql.NullString{String: employee.Nip},
		Birthplace:  sql.NullString{String: employee.Birthplace},
		Birthdate:   sql.NullTime{Time: birthdate},
		Age:         sql.NullInt32{Int32: employee.Age},
		Address:     sql.NullString{String: employee.Address},
		Religion:    sql.NullString{String: employee.Religion},
		Gender:      sql.NullString{String: employee.Gender},
		PhoneNumber: sql.NullString{String: employee.PhoneNumber},
		Email:       sql.NullString{String: employee.Email},
		CreatedAt:   sql.NullTime{Time: timeNow},
		CreatedBy:   sql.NullInt64{Int64: context.UserLoginID},
		UpdatedAt:   sql.NullTime{Time: timeNow},
		UpdatedBy:   sql.NullInt64{Int64: context.UserLoginID},
	}

	err = t.employeeRepo.Add(nil, tx, &employeeModel)
	if err.Err != nil {
		return
	}

	err = util.GenerateNonError()
	return
}
