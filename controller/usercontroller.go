package controller

import (
	"net/http"
	"encoding/json"
	"workspace/model"
	"workspace/service"
)

type UserController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(newUserService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: newUserService,
	}
}

func (this *UserControllerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
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
	saveResponse, err := this.userService.SignUp(req)
	json.NewEncoder(w).Encode(saveResponse)
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
}