package services

import (
	"github.com/lambertse/cquan_go_webapp/db"
	"github.com/lambertse/cquan_go_webapp/models"
	"gorm.io/gorm"
)

type UserService struct {
  DB *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{DB: db.DB}
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
  var users []*models.User
  if err := s.DB.Find(&users).Error; err != nil {
    return nil, err
  }
  return users, nil
}


