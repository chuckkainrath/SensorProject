package repository

import "SensorProject/models"

type IUserRepo interface {
	GetUser(username string) (*models.User, error)
}

type userRepo struct {
}

func NewUserRepo() IUserRepo {
	return userRepo{}
}

func (u userRepo) GetUser(username string) (*models.User, error) {
	var user *models.User
	result := DB().Where("user_name = ?", username).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
