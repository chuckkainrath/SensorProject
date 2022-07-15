package service

import (
	"time"

	"github.com/chuckkainrath/SensorProject/middleware/errors"
	"github.com/chuckkainrath/SensorProject/models"
	"github.com/chuckkainrath/SensorProject/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserToken(username, password string) (*string, *errors.AppError)
}

type userService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{UserRepo: userRepo}
}

func (u userService) GetUserToken(username, password string) (*string, *errors.AppError) {
	user, err := u.UserRepo.GetUser(username)
	if err != nil {
		return nil, err
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if errPassword != nil {
		return nil, errors.NewBadRequestError("Username/Password combination not found")
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	tk := &models.Token{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, tknErr := token.SignedString([]byte("randomsecretstring"))
	if tknErr != nil {
		return nil, errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return &tokenString, nil
}
