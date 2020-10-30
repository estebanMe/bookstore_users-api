package services

import (
	"github.com/estebanMe/bookstore_users-api/domain/users"
	"github.com/estebanMe/bookstore_users-api/utils/criptoutils"
	"github.com/estebanMe/bookstore_users-api/utils/dateutils"
	"github.com/estebanMe/bookstore_users-api/utils/errors"
)

//GetUser get user by id
func GetUser(userID int64) (*users.User, *errors.RestErr) {

	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

//CreateUser Create user services
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

    user.Status = users.StatusActive
	user.DateCreated = dateutils.GetNowDBFormat()
	user.Password = criptoutils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

//UpdateUser UpdateUser
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil

}

//DeleteUser delete user service
func DeleteUser(userID int64) *errors.RestErr {
 user := &users.User{ID: userID}
 return user.Delete()
}

//Search service of find record by status value 
func Search(status string)  (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)

	
}
