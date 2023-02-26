package models

type User struct {
	Id        int    `json:"id"`
	Age       int    `json:"age"`
	FirstName string `json:"first_Name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
