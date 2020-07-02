package users

import (
	"bookstore_users-api/datasource/mysql/users_db"
	"bookstore_users-api/util/date_utils"
	"bookstore_users-api/util/errors"
	"bookstore_users-api/util/msql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name,last_name,email,date_create) VALUES(?,?,?,?);"
	queryGetUser    = "SELECT id,first_name,last_name,email,date_create FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name = ?,last_name=?,email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

//Obtengo el Id del Usuario
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return msql_utils.ParseError(getErr)
	}
	return nil
}

//Save usuario
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return msql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return msql_utils.ParseError(saveErr)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return msql_utils.ParseError(err)
	}
	return nil

}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	if _, err = stmt.Exec(user.Id); err != nil {
		return msql_utils.ParseError(err)
	}

	return nil
}
