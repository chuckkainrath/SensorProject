package service

import (
	"SensorProject/models"
	"SensorProject/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type IUserService interface {
	GetUserToken(username, password string) (*string, error)
}

type userService struct {
	UserRepo repository.IUserRepo
}

func NewUserService() IUserService {
	return userService{UserRepo: repository.NewUserRepo()}
}

func (u userService) GetUserToken(username, password string) (*string, error) {
	user, err := u.UserRepo.GetUser(username)
	if err != nil {
		// TODO: Figure out error response
		return nil, err
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if errPassword != nil {
		// TODO: Figure out error response
		return nil, errPassword
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	tk := &models.Token{
		UserID:   user.ID,
		Username: user.Username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		// TODO: Figure out error response
		return nil, err
	}
	return &tokenString, nil
}
