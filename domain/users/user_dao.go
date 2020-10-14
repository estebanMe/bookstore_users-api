package users

import (
	"fmt"
	"strings"

	"github.com/estebanMe/bookstore_users-api/datasource/mysql/usersdb"
	"github.com/estebanMe/bookstore_users-api/utils/dateutils"
	"github.com/estebanMe/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser = ("Insert INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);")
)

var (
	usersDB = make(map[int64]*User)
)

//Get get record by userId
func (user *User) Get() *errors.RestErr {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("El map con las get querys", usersDB)
	result := usersDB[user.ID]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
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
	
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail){
			return errors.NewBadRequestError("email %s already exists")
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when triying to save user: %s", err.Error()),
		)
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
