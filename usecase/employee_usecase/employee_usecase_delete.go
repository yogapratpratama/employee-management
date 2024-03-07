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

func (t employeeUsecase) Delete(_ context.Context, context domain.ContextModel, employee *domain.EmployeeRequest) (err util.ErrorModel) {
	var (
		fileName      = "employee_usecase_delete.go"
		funcName      = "Delete"
		db            = serverconfig.ServerAttribute.DBConnection
		timeNow       = time.Now()
		employeeModel model.EmployeeModel
		employeeOnDB  model.EmployeeModel
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

	employeeModel = model.EmployeeModel{
		ID:        sql.NullInt64{Int64: employee.ID},
		DeletedAt: sql.NullTime{Time: timeNow},
		UpdatedBy: sql.NullInt64{Int64: context.UserLoginID},
		UpdatedAt: sql.NullTime{Time: timeNow},
	}

	employeeOnDB, err = t.employeeRepo.GetByID(nil, db, employeeModel.ID.Int64)
	if err.Err != nil {
		return
	}

	if employeeOnDB.ID.Int64 < 1 {
		err = util.GenerateNotFoundError(fileName, funcName)
		return
	}

	err = t.employeeRepo.Delete(nil, tx, employeeModel)
	if err.Err != nil {
		return
	}

	err = util.GenerateNonError()
	return
}
