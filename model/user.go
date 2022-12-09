package model

import (
	"time"
)

type User struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	Password string `json:"password"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func (req *UserRequest) ToUser() User {
	user := User{ID: 15, Email: req.Email, Password: req.Password, FirstName: req.FirstName, LastName: req.LastName}
	return user 
}