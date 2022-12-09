package service

import (
	"workspace/model"
	"workspace/repository"
)

type UserService interface {
	SignUp(userRequest model.UserRequest) (model.SaveResponse, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(newUserRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: newUserRepository,
	}
}

func (self *UserServiceImpl) SignUp(userRequest model.UserRequest) (model.SaveResponse, error) {
	var user model.User
	user = userRequest.ToUser()
	id, err := self.userRepository.Save(user)
	if err != nil {
		return model.SaveResponse{}, err
	}
	return model.SaveResponse{id}, nil
}