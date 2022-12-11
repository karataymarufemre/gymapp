package repository

import (
	"database/sql"
	"fmt"
	"workspace/model"
	"workspace/service/authservice"
)

type UserRepository interface {
	Save(user model.User) error
	Get(request model.UserLoginRequest) (model.UserDTO, error)
}

type UserRepositoryImpl struct {
	db          *sql.DB
	authService authservice.AuthService
}

func NewUserRepository(database *sql.DB, authService authservice.AuthService) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db:          database,
		authService: authService,
	}
}

func (r *UserRepositoryImpl) Save(user model.User) error {
	sqlStatement := InsertQuery("user", "email, password, first_name, last_name")
	_, err := r.db.Exec(
		sqlStatement,
		user.Email, user.Password, user.FirstName, user.LastName,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Get(request model.UserLoginRequest) (model.UserDTO, error) {
	var userDTO model.UserDTO
	var password string
	sqlStatement := "SELECT id, password, email, first_name, last_name FROM user WHERE email=?"
	err := r.db.QueryRow(sqlStatement, request.Email).Scan(&userDTO.ID, &password, &userDTO.Email, &userDTO.FirstName, &userDTO.LastName)
	if err != nil {
		return userDTO, fmt.Errorf("could not find user. err: %w", err)
	}
	isPassword := r.authService.CheckPasswordHash(request.Password, password)
	if !isPassword {
		return userDTO, fmt.Errorf("password is wrong")
	}

	return userDTO, nil
}
