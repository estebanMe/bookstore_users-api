package users

import (
	"fmt"

	"github.com/estebanMe/bookstore_users-api/datasource/mysql/usersdb"
	"github.com/estebanMe/bookstore_users-api/utils/dateutils"
	"github.com/estebanMe/bookstore_users-api/utils/errors"
	"github.com/estebanMe/bookstore_users-api/utils/mysqlutils"
)

const (
	queryInsertUser  = "Insert INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

//Get get record by userId
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysqlutils.ParseError(getErr)
	}

	return nil
}

//Save save user and return nil value
func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = dateutils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)
	
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {

		return errors.NewInternalServerError(
			fmt.Sprintf("error when triying to save user: %s", err.Error()),
		)
	}

	user.ID = userID

	return nil
}
