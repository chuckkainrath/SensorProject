package controllers

import (
	"SensorProject/dtos"
	"SensorProject/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type IUserController interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	UserService service.IUserService
}

func NewUserController() IUserController {
	return userController{UserService: service.NewUserService()}
}

func (u userController) Login(w http.ResponseWriter, r *http.Request) {
	user := &dtos.UserDto{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// TODO: Invalid Request Response
		return
	}
	token, err := u.UserService.GetUserToken(user.UserName, user.Password)
	if err != nil {
		// TODO: Could not login
		return
	}
	// TODO: Token response (modify line below)
	fmt.Fprintf(w, "Bearer %s\n", *token)
}
