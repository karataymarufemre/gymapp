package service

import (
	"errors"
	"fmt"
	"workspace/internal/model"
	"workspace/internal/repository"
	"workspace/internal/service/authservice"
)

type UserService interface {
	SignUp(userRequest model.UserRequest) (model.SaveResponse, error)
	SignIn(request model.UserLoginRequest) (string, error)
	GetAccessToken(jwtClaims authservice.JWTClaims) (string, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
	authService    authservice.AuthService
}

func NewUserService(newUserRepository repository.UserRepository, newAuthService authservice.AuthService) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: newUserRepository,
		authService:    newAuthService,
	}
}

func (s *UserServiceImpl) SignUp(userRequest model.UserRequest) (model.SaveResponse, error) {
	user := userRequest.ToUser()
	user.Password, _ = s.authService.HashPassword(user.Password)
	err := s.userRepository.Save(user)
	if err != nil {
		return model.SaveResponse{Success: false}, err
	}
	return model.SaveResponse{Success: true}, nil
}

func (s *UserServiceImpl) SignIn(request model.UserLoginRequest) (string, error) {
	user, err := s.userRepository.Get(request)
	if err != nil {
		return "", err
	}
	token, err := s.authService.GenerateJWT(user.ID, 129600, true)
	if err != nil {
		return "", fmt.Errorf("could not generate JWT token. err: %w", err)
	}
	return token, nil
}

func (s *UserServiceImpl) GetAccessToken(jwtClaims authservice.JWTClaims) (string, error) {
	if !jwtClaims.IsLongToken {
		return "", errors.New("This token is not a type of refresh token.")
	}
	token, err := s.authService.GenerateJWT(jwtClaims.UserId, 20, false)
	if err != nil {
		return "", fmt.Errorf("could not generate JWT token. err: %w", err)
	}
	return token, nil
}
