package controllers

import (
	"SensorProject/models"
	"SensorProject/repository"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Err string
}

var db = repository.DB()

func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindUser(user.Username, user.HashedPassword)
	json.NewEncoder(w).Encode(resp)
}

func FindUser(username, password string) map[string]interface{} {
	user := &models.User{}

	if err := db.Where("Username = ?", username).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Username not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if errPassword != nil && errPassword == bcrypt.ErrMismatchedHashAndPassword { //Password does not match
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Token{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}
