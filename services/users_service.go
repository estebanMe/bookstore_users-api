package services

import (
	"github.com/estebanMe/bookstore_users-api/domain/users"
	"github.com/estebanMe/bookstore_users-api/utils/errors"
)

//CreateUser Create user services
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}