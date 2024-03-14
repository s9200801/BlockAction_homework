package service

import "gorm.io/gorm"

type baseService struct {
	DB *gorm.DB
}

func newBaseService(db *gorm.DB) baseService {
	return baseService{
		DB:     db,
	}
}