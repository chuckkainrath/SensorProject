package controllers

import (
	"SensorProject/dtos"
	"SensorProject/middleware"
	"SensorProject/service"
	"net/http"
)

type userLoginHandler struct {
	UserService service.UserService
}

func NewUserLoginHandler(userService service.UserService) http.Handler {
	return &userLoginHandler{UserService: userService}
}

func (u *userLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.Login(w, r)
}

func (u *userLoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	user := **middleware.GetRequestBody(r).(**dtos.UserDto)

	token, err := u.UserService.GetUserToken(user.UserName, user.Password)
	if err != nil {
		middleware.AddResultToContext(r, *err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, token, middleware.OutputDataKey)
}
