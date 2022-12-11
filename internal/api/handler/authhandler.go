package handler

import (
	"net/http"
	"workspace/internal/api/controller"
	"workspace/internal/constants/urlconstants"
)

type AuthHandler interface {
	Handle(mux *http.ServeMux)
}

type AuthHandlerImpl struct {
	userController controller.UserController
}

func NewAuthHandler(userController controller.UserController) *AuthHandlerImpl {
	return &AuthHandlerImpl{
		userController: userController,
	}
}

func (a *AuthHandlerImpl) Handle(mux *http.ServeMux) {
	mux.HandleFunc(urlconstants.AUTH_SIGN_UP, a.userController.SignUp)
	mux.HandleFunc(urlconstants.AUTH_SIGN_IN, a.userController.SignIn)
	mux.HandleFunc(urlconstants.AUTH_REFRESH_TOKEN, a.userController.GetAccessToken)
}
