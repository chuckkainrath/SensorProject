package repository

import "gorm.io/gorm"

type repositoryPostgreSQL struct {
	db *gorm.DB
}
