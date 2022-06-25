package repository

import "gorm.io/gorm"

type RepositoryPostgreSQL struct {
	db *gorm.DB
}
