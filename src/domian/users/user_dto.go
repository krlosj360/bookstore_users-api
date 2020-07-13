package users

import (
	"bookstore_users-api/src/util/errors"
	"encoding/json"
	"strings"
	"time"
)

const (
	StatusActive = "ACTIVE"
)

// first create a type alias
type JsonBirthDate time.Time

type User struct {
	Id              int64  `json:"id"`
	Identification  string `json:"identification"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	BirthDate       string `json:"birthDate"`
	SenescytId      string `json:"senescyt_Id"`
	UniversityTitle string `json:"university_title"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Agree           bool   `json:"agree"`
	Status          string `json:"state"`
	DateCreated     string `json:"date_created"`
	DateUpdated     string `json:"date_updated"`
	RoleId          string `json:"role_id"`
}

type Users []User

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	//user.BirthDate = user.BirthDate.Format("2006-01-02")
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}

// imeplement Marshaler und Unmarshalere interface
func (j *JsonBirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonBirthDate(t)
	return nil
}

func (j JsonBirthDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// Maybe a Format function for printing your date
func (j JsonBirthDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
