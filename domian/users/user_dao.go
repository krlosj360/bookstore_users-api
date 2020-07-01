package users

import (
	"bookstore_users-api/datasource/mysql/users_db"
	"bookstore_users-api/util/date_utils"
	"bookstore_users-api/util/errors"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
)

//Obtengo el Id del Usuario
func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %d already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	//TODO:Llama a funcion de crear fecha
	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
