package models

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"-"`
	Phone     string `json:"phone"`
}
