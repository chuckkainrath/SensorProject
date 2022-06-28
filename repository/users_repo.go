package repository

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(username string) (*models.User, *errors.AppError)
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepositoryDB(db *gorm.DB) userRepository {
	return userRepository{db}
}

func (u userRepository) GetUser(username string) (*models.User, *errors.AppError) {
	var user *models.User
	result := DB().Where("user_name = ?", username).First(user)
	if result.Error != nil {
		return nil, errors.NewNotFoundError("User not found")
	}
	return user, nil
}
