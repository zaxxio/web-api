package service

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"web/model"
)

// UserService handles all user-related database operations.
type UserService struct {
	DB *gorm.DB
}

// NewUserService is the Fx constructor that provides UserService.
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// CreateUser inserts a new user into the database.
func (s *UserService) CreateUser(u *model.User) error {
	return s.DB.Create(u).Error
}

// UpdateUser updates user information in the database.
func (s *UserService) UpdateUser(u *model.User) error {
	return s.DB.Save(u).Error
}

// DeleteUser removes a user from the database.
func (s *UserService) DeleteUser(u *model.User) error {
	return s.DB.Delete(u).Error
}

// GetUsers returns all users in the database.
func (s *UserService) GetUsers() ([]model.User, error) {
	var users []model.User
	err := s.DB.Find(&users).Error
	return users, err
}

// GetUserByEmail finds a user by their email address.
func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	err := s.DB.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// UserModule is the Fx module that provides this service.
var UserModule = fx.Module(
	"service",
	fx.Provide(NewUserService),
)
