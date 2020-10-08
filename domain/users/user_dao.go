package users

import (
	"fmt"

	"github.com/estebanMe/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

//Get get record by userId
func (user *User) Get() *errors.RestErr {
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
	current := usersDB[user.ID]
	fmt.Println("The current value before save method", current)
	fmt.Println("The user email value from model", user.Email)
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("user %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}

	usersDB[user.ID] = user

	return nil
}
