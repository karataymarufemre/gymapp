package repository

import (
	"database/sql"
	"workspace/model"
)

type UserRepository interface {
	Save(user model.User) (int64, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: database,
	}
}

func (self *UserRepositoryImpl) Save(user model.User) (int64, error) {
	sqlStatement := `
		INSERT INTO users (age, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)`
	return 998, nil
}
