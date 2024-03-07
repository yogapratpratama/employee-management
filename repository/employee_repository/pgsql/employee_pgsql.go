package pgsql

import (
	"EmployeeManagementApp/domain"
	"EmployeeManagementApp/repository/employee_repository/model"
	util_repo "EmployeeManagementApp/repository/util"
	"EmployeeManagementApp/util"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type pgsqlEmployeeRepository struct {
	FileName  string
	TableName string
	Conn      *sql.DB
}

func NewPgsqlEmployeeRepository(conn *sql.DB) domain.EmployeeRepository {
	return &pgsqlEmployeeRepository{
		FileName:  "employee_pgsql.go",
		TableName: "employee",
		Conn:      conn,
	}
}

func (p pgsqlEmployeeRepository) Add(_ context.Context, tx *sql.Tx, employee *model.EmployeeModel) util.ErrorModel {
	var (
		funcName = "Add"
		query    string
		param    []interface{}
	)

	query = fmt.Sprintf(`
		INSERT INTO %s 
			(
			 name, nip, birthplace, 
			 birthdate, age, address, 
			 religion, gender, phone_number, 
			 email, created_at, updated_at,
			 created_by, updated_by
			) 
		VALUES 
		($1, $2, $3, 
		$4, $5, $6, 
		$7, $8, $9, 
		$10, $11, $12,
		$13, $14) 
		RETURNING id`,
		p.TableName)

	param = append(param,
		employee.Name.String, employee.Nip.String, employee.Birthplace.String,
		employee.Birthdate.Time, employee.Age.Int32, employee.Address.String,
		employee.Religion.String, employee.Gender.String, employee.PhoneNumber.String,
		employee.Email.String, employee.CreatedAt.Time, employee.UpdatedAt.Time,
		employee.CreatedBy.Int64, employee.UpdatedBy.Int64)

	result := tx.QueryRow(query, param...)
	errs := result.Scan(&employee.ID)
	if errs != nil && errs.Error() != sql.ErrNoRows.Error() {
		return util.GenerateInternalDBServerError(p.FileName, funcName, errs)
	}

	return util.GenerateNonError()
}

func (p pgsqlEmployeeRepository) Update(_ context.Context, tx *sql.Tx, employee model.EmployeeModel) util.ErrorModel {
	var (
		funcName = "Update"
		query    string
	)

	query = fmt.Sprintf(`
		UPDATE %s 
		SET name = $1, nip = $2, birthplace = $3,
			birthdate = $4, age = $5, address = $6, 
			religion = $7, gender = $8, phone_number = $9, 
			email = $10, updated_at = $11, updated_by = $12 
		WHERE id = $13 AND deleted = FALSE `,
		p.TableName)

	params := []interface{}{
		employee.Name.String, employee.Nip.String, employee.Birthplace.String,
		employee.Birthdate.Time, employee.Age.Int32, employee.Address.String,
		employee.Religion.String, employee.Gender.String, employee.PhoneNumber.String,
		employee.Email.String, employee.UpdatedAt.Time, employee.UpdatedBy.Int64,
		employee.ID.Int64}

	return util_repo.UpdateRow(nil, tx, query, params, p.FileName, funcName)
}

func (p pgsqlEmployeeRepository) Delete(_ context.Context, tx *sql.Tx, employee model.EmployeeModel) util.ErrorModel {
	var (
		funcName = "Delete"
		query    string
	)

	query = fmt.Sprintf(`
		UPDATE %s 
		SET deleted = TRUE, updated_at = $1, updated_by = $2, deleted_at = $3
		WHERE id = $4 AND deleted = FALSE `,
		p.TableName)

	params := []interface{}{
		employee.UpdatedAt.Time, employee.UpdatedBy.Int64, employee.DeletedAt.Time,
		employee.ID.Int64}
	return util_repo.UpdateRow(nil, tx, query, params, p.FileName, funcName)
}

func (p pgsqlEmployeeRepository) Fetch(_ context.Context, db *sql.DB, param util_repo.GetListParameterModel) (output []model.EmployeeModel, err util.ErrorModel) {
	var (
		funcName = "Fetch"
		query    string
		params   []interface{}
	)

	query = fmt.Sprintf(`
		SELECT 
		    id, name, nip, 
		    email, phone_number 
		FROM %s
		WHERE deleted = FALSE `,
		p.TableName)

	if len(param.Filter) > 0 {
		for i, itemFilter := range param.Filter {
			switch itemFilter.Key.String {
			case "name":
				param.Filter[i].Key.String = "name"
				if param.Filter[i].Operator.String == "like" {
					query += " AND LOWER(name) like '%" + param.Filter[i].Value.String + "%'"
				} else if param.Filter[i].Operator.String == "eq" {
					query += " AND name = '" + param.Filter[i].Value.String + "'"
				} else {
					err = util.GenerateInternalDBServerError(p.FileName, funcName, errors.New("name only eq and like"))
					return
				}
			case "nip":
				param.Filter[i].Key.String = "nip"
				if param.Filter[i].Operator.String == "eq" {
					query += " AND nip = '" + param.Filter[i].Value.String + "'"
				} else {
					err = util.GenerateInternalDBServerError(p.FileName, funcName, errors.New("nip only eq"))
					return
				}
			case "email":
				param.Filter[i].Key.String = "email"
				if param.Filter[i].Operator.String == "eq" {
					query += " AND email = " + param.Filter[i].Value.String
				} else {
					err = util.GenerateInternalDBServerError(p.FileName, funcName, errors.New("email only eq"))
					return
				}
			default:
			}
		}
	}

	if len(param.Order) > 0 {
		query += " ORDER BY "
		for i, itemOrder := range param.Order {
			query += itemOrder.Order.String
			if len(param.Order)-(i+1) > 0 {
				query += ", "
			}
		}
	}

	query += fmt.Sprintf(` LIMIT $1 OFFSET $2`)
	params = []interface{}{param.Limit.Int64, util_repo.CountOffset(int(param.Page.Int64), int(param.Limit.Int64))}

	rows, errs := db.Query(query, params...)
	if errs != nil && errs.Error() != sql.ErrNoRows.Error() {
		err = util.GenerateInternalDBServerError(p.FileName, funcName, errs)
		return
	}

	var result []interface{}
	result, err = util_repo.GetRows(rows, func(rws *sql.Rows) (interface{}, error) {
		var temp model.EmployeeModel
		errs1 := rws.Scan(
			&temp.ID, &temp.Name, &temp.Nip,
			&temp.Email, &temp.PhoneNumber)
		return temp, errs1
	})

	if len(result) > 0 {
		for _, itemResult := range result {
			output = append(output, itemResult.(model.EmployeeModel))
		}
	}

	err = util.GenerateNonError()
	return
}

func (p pgsqlEmployeeRepository) GetByID(_ context.Context, db *sql.DB, id int64) (output model.EmployeeModel, err util.ErrorModel) {
	var (
		funcName = "GetByID"
		query    string
	)

	query = fmt.Sprintf(`
		SELECT 
			id, name, nip, 
			birthplace, birthdate, age, 
			address, religion, gender,
			phone_number, email, updated_at,
			created_at
		FROM %s
		WHERE id = $1 AND deleted = FALSE `,
		p.TableName)

	params := []interface{}{id}
	results := db.QueryRow(query, params...)
	dbErr := results.Scan(
		&output.ID, &output.Name, &output.Nip,
		&output.Birthplace, &output.Birthdate, &output.Age,
		&output.Address, &output.Religion, &output.Gender,
		&output.PhoneNumber, &output.Email, &output.UpdatedAt,
		&output.CreatedAt)

	if dbErr != nil && dbErr.Error() != sql.ErrNoRows.Error() {
		err = util.GenerateInternalDBServerError(p.FileName, funcName, dbErr)
		return
	}

	err = util.GenerateNonError()
	return
}
