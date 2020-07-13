package model

type User struct {
	Identification  string `json:"identification"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	BirthDate       string `json:"birthDate"`
	SenescytId      string `json:"senescyt_Id"`
	UniversityTitle string `json:"university_title"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Agree           bool   `json:"agree"`
	RoleId          string `json:"role_id"`
}
