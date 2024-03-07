package pgsql

import (
	"EmployeeManagementApp/domain"
	"EmployeeManagementApp/util"
	"context"
	"database/sql"
	"fmt"
)

type pgsqlUserRepository struct {
	FileName  string
	TableName string
	Conn      *sql.DB
}

func NewPgsqlUserRepository(conn *sql.DB) domain.UserRepository {
	return &pgsqlUserRepository{
		FileName:  "user_pgsql.go",
		TableName: "user",
		Conn:      conn,
	}
}

func (p pgsqlUserRepository) Store(_ context.Context, tx *sql.Tx, users *domain.User) util.ErrorModel {
	funcName := "Store"
	query := fmt.Sprintf(`
		INSERT INTO "%s" 
		(
		 username, password, created_at, updated_at
		) 
		VALUES 
		(
		    $1, $2, $3, $4
		) returning id `,
		p.TableName)

	params := []interface{}{
		users.Username, users.Password, users.CreatedAt, users.UpdatedAt}
	result := tx.QueryRow(query, params...)
	errs := result.Scan(&users.ID)
	if errs != nil && errs.Error() != sql.ErrNoRows.Error() {
		return util.GenerateInternalDBServerError(p.FileName, funcName, errs)
	}

	return util.GenerateNonError()
}

func (p pgsqlUserRepository) GetByUsername(_ context.Context, db *sql.DB, username string) (output domain.User, err util.ErrorModel) {
	funcName := "GetByUsername"
	query := fmt.Sprintf(` 
		SELECT id, username 
		FROM "%s" 
		WHERE username = $1 AND deleted = FALSE `,
		p.TableName)

	params := []interface{}{username}
	results := db.QueryRow(query, params...)
	dbErr := results.Scan(
		&output.ID, &output.Username)

	if dbErr != nil && dbErr.Error() != sql.ErrNoRows.Error() {
		err = util.GenerateInternalDBServerError(p.FileName, funcName, dbErr)
		return
	}

	err = util.GenerateNonError()
	return
}

func (p pgsqlUserRepository) GetLogin(_ context.Context, db *sql.DB, users *domain.User) (output domain.User, err util.ErrorModel) {
	funcName := "GetLogin"
	query := fmt.Sprintf(` 
		SELECT id, password
		FROM "%s"
		WHERE username = $1 AND deleted = FALSE`,
		p.TableName)

	params := []interface{}{users.Username}
	results := db.QueryRow(query, params...)
	dbErr := results.Scan(&output.ID, &output.Password)
	if dbErr != nil && dbErr.Error() != sql.ErrNoRows.Error() {
		err = util.GenerateInternalDBServerError(p.FileName, funcName, dbErr)
		return
	}

	err = util.GenerateNonError()
	return
}
