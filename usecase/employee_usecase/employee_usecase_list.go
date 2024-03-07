package employee_usecase

import (
	"EmployeeManagementApp/app/serverconfig"
	"EmployeeManagementApp/domain"
	"EmployeeManagementApp/repository/employee_repository/model"
	"EmployeeManagementApp/repository/util"
	util2 "EmployeeManagementApp/util"
	"context"
	"database/sql"
	"errors"
	"strings"
)

func (t employeeUsecase) Fetch(_ context.Context, employee *domain.EmployeeRequest) (output []domain.ListEmployeeResponse, err util2.ErrorModel) {
	var (
		fileName      = "employee_usecase_list.go"
		funcName      = "Fetch"
		db            = serverconfig.ServerAttribute.DBConnection
		searchByParam util.GetListParameterModel
		listFilter    []util.ListFilter
		listOrder     []util.ListOrder
		listDataDB    []model.EmployeeModel
	)

	if employee.Filter != "" {
		filter := strings.Split(employee.Filter, ",")
		for _, itemFilter := range filter {
			filterComponent := strings.Split(itemFilter, " ")
			if len(filterComponent) != 3 {
				err = util2.GenerateUnknownServerError(fileName, funcName, errors.New("wrong format"))
				return
			}

			if filterComponent[1] != "eq" && filterComponent[1] != "like" {
				err = util2.GenerateUnknownServerError(fileName, funcName, errors.New("wrong format must eq or like"))
				return
			}

			listFilter = append(listFilter, util.ListFilter{
				Key:      sql.NullString{String: filterComponent[0]},
				Operator: sql.NullString{String: filterComponent[1]},
				Value:    sql.NullString{String: filterComponent[2]},
			})
		}
	}

	if employee.Order != "" {
		order := strings.Split(employee.Order, ",")
		for _, itemOrder := range order {
			listOrder = append(listOrder, util.ListOrder{Order: sql.NullString{String: itemOrder}})
		}
	}

	searchByParam = util.GetListParameterModel{
		Page:   sql.NullInt64{Int64: employee.Page},
		Limit:  sql.NullInt64{Int64: employee.Limit},
		Filter: listFilter,
		Order:  listOrder,
	}

	listDataDB, err = t.employeeRepo.Fetch(nil, db, searchByParam)
	if err.Err != nil {
		return
	}

	if len(listDataDB) > 0 {
		for _, itemListDataDB := range listDataDB {
			output = append(output, domain.ListEmployeeResponse{
				ID:          itemListDataDB.ID.Int64,
				Name:        itemListDataDB.Name.String,
				Nip:         itemListDataDB.Nip.String,
				Email:       itemListDataDB.Email.String,
				PhoneNumber: itemListDataDB.PhoneNumber.String,
			})
		}
	}

	err = util2.GenerateNonError()
	return
}
