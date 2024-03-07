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

func (t employeeUsecase) Update(_ context.Context, context domain.ContextModel, employee *domain.EmployeeUpdateRequest) (err util.ErrorModel) {
	var (
		fileName      = "employee_usecase_update.go"
		funcName      = "Update"
		db            = serverconfig.ServerAttribute.DBConnection
		timeNow       = time.Now()
		employeeModel model.EmployeeModel
		employeeOnDB  model.EmployeeModel
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
		ID:          sql.NullInt64{Int64: employee.ID},
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
		UpdatedAt:   sql.NullTime{Time: timeNow},
		UpdatedBy:   sql.NullInt64{Int64: context.UserLoginID},
	}

	employeeOnDB, err = t.employeeRepo.GetByID(nil, db, employeeModel.ID.Int64)
	if err.Err != nil {
		return
	}

	if employeeOnDB.ID.Int64 < 1 {
		err = util.GenerateNotFoundError(fileName, funcName)
		return
	}

	err = t.employeeRepo.Update(nil, tx, employeeModel)
	if err.Err != nil {
		return
	}

	err = util.GenerateNonError()
	return
}
