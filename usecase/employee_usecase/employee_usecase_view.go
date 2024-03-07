package employee_usecase

import (
	"EmployeeManagementApp/app/serverconfig"
	"EmployeeManagementApp/domain"
	"EmployeeManagementApp/repository/employee_repository/model"
	"EmployeeManagementApp/util"
	"context"
)

func (t employeeUsecase) GetByID(_ context.Context, id int64) (output domain.ViewEmployeeResponse, err util.ErrorModel) {
	var (
		fileName = "employee_usecase_view.go"
		funcName = "GetByID"
		viewDB   model.EmployeeModel
		db       = serverconfig.ServerAttribute.DBConnection
	)

	viewDB, err = t.employeeRepo.GetByID(nil, db, id)
	if err.Err != nil {
		return
	}

	if viewDB.ID.Int64 < 1 {
		err = util.GenerateNotFoundError(fileName, funcName)
		return
	}

	output = domain.ViewEmployeeResponse{
		ID:          viewDB.ID.Int64,
		Name:        viewDB.Name.String,
		Nip:         viewDB.Nip.String,
		Email:       viewDB.Email.String,
		PhoneNumber: viewDB.PhoneNumber.String,
		Birthplace:  viewDB.Birthplace.String,
		Birthdate:   viewDB.Birthdate.Time,
		Age:         viewDB.Age.Int32,
		Address:     viewDB.Address.String,
		Religion:    viewDB.Religion.String,
		Gender:      viewDB.Gender.String,
		CreatedAt:   viewDB.CreatedAt.Time,
		UpdatedAt:   viewDB.UpdatedAt.Time,
	}

	err = util.GenerateNonError()
	return
}
