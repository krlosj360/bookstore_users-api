package users

import (
	"bookstore_users-api/datasource/mysql/users_db"
	"bookstore_users-api/util/errors"
	"bookstore_users-api/util/msql_utils"
	"fmt"
	"log"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name,last_name,email,password,status,date_create) VALUES(?,?,?,?,?,?);"
	queryGetUser          = "SELECT id,first_name,last_name,email,date_create,status FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name = ?,last_name=?,email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id,first_name,last_name,email,date_create,status FROM users WHERE status=?;"
)

//Obtengo el Id del Usuario
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
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
	//user.DateCreated = date_utils.GetNowDBFormat()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Status, user.DateCreated)
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

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())

	}

	defer rows.Close()
	results := make([]User, 0)

	for rows.Next() {
		var user User

		//SELECT id,first_name,last_name,email,date_create,status
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			log.Println("Aqui se callo -")
			return nil, msql_utils.ParseError(err)

		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user matching status %s", status))
	}
	return results, nil
}
