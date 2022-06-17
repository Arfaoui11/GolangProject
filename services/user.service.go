package services

import "CrudGolang/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(*string) error
}
