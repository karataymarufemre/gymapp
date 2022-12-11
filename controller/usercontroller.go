package controller

import (
	"encoding/json"
	"net/http"
	"workspace/model"
	"workspace/service"
	"workspace/service/authservice"
)

type UserController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	GetAccessToken(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(newUserService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: newUserService,
	}
}

func (c *UserControllerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	var req model.UserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	saveResponse, err := c.userService.SignUp(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(saveResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (c *UserControllerImpl) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	var req model.UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := c.userService.SignIn(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := model.TokenResponse{Token: resp}
	json.NewEncoder(w).Encode(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (c *UserControllerImpl) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	claims := r.Context().Value("claims").(authservice.JWTClaims)
	resp, err := c.userService.GetAccessToken(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := model.TokenResponse{Token: resp}
	json.NewEncoder(w).Encode(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
