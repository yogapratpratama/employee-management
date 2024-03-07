package model

import "database/sql"

type EmployeeModel struct {
	ID          sql.NullInt64
	Name        sql.NullString
	Nip         sql.NullString
	Birthplace  sql.NullString
	Birthdate   sql.NullTime
	Age         sql.NullInt32
	Address     sql.NullString
	Religion    sql.NullString
	Gender      sql.NullString
	PhoneNumber sql.NullString
	Email       sql.NullString
	CreatedAt   sql.NullTime
	CreatedBy   sql.NullInt64
	UpdatedAt   sql.NullTime
	UpdatedBy   sql.NullInt64
	DeletedAt   sql.NullTime
}
