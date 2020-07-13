package users

import (
	"bookstore_users-api/src/datasource/mysql/users_db"
	"bookstore_users-api/src/logger"
	"bookstore_users-api/src/util/errors"
	"bookstore_users-api/src/util/msql_utils"
	"fmt"
	"strings"
)

const (
	queryInsertUser             = "INSERT INTO users(identification,firstname,lastname,birthdate,senescyt_id,university_title,email,password,agree,state,date_created,date_updated,role_id)  VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);"
	queryGetUser                = "SELECT id,firstname,lastname,email,date_created,state FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET firstname = ?,lastname=?,email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus       = "SELECT id,firstname,last_name,email,date_created FROM users WHERE state=?;"
	queryFindByEmailAndPassword = "SELECT id, firstname, last_name, email, date_created, state FROM users WHERE email=? AND password=? AND state=?"
)

//Obtengo el Id del Usuario
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to prepare get user by id", getErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

//Save usuario
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	//user.DateCreated = date_utils.GetNowDBFormat()
	insertResult, saveErr := stmt.Exec(user.Identification, user.FirstName, user.LastName, user.BirthDate, user.SenescytId, user.UniversityTitle, user.Email, user.Password, user.Agree, user.Status, user.DateCreated, user.DateUpdated, user.RoleId)
	if saveErr != nil {
		fmt.Printf(saveErr.Error())
		logger.Error("error when trying to prepare save user", err)
		return errors.NewInternalServerError("database error")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error trying prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		errors.NewInternalServerError("database error")
	}
	return nil

}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying delete user", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) FindByStatus(status string) (Users, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)

	if err != nil {
		logger.Error("error when trying to find user by status ", err)
		return nil, errors.NewInternalServerError("database error")

	}

	defer rows.Close()
	results := make(Users, 0)

	for rows.Next() {
		var user User

		//SELECT id,first_name,last_name,email,date_create,status
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")

		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user matching status %s", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), msql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", getErr)
		return errors.NewInternalServerError("error when tying to find user")
	}
	return nil
}
