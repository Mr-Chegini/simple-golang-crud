package service

import (
	"errors"

	"github.com/Mr-Chegini/simple-golang-crud/models"
	"github.com/Mr-Chegini/simple-golang-crud/repository"
	"github.com/Mr-Chegini/simple-golang-crud/utils"
)

// UserServiceInterface defines business operations for users.
type UserServiceInterface interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	ListUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

// UserService implements UserServiceInterface.
type UserService struct {
	userRepo repository.UserRepositoryInterface
}

// NewUserService creates a new UserService.
func NewUserService(userRepo repository.UserRepositoryInterface) *UserService {
	return &UserService{userRepo: userRepo}
}

// CreateUser validates and creates a new user.
func (s *UserService) CreateUser(user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	if !utils.IsNonEmpty(user.Name) {
		return errors.New("user name is required")
	}

	if !utils.IsValidEmail(user.Email) {
		return errors.New("valid user email is required")
	}

	return s.userRepo.Create(user)
}

// GetUserByID returns a single user.
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	if id == 0 {
		return nil, errors.New("invalid user id")
	}
	return s.userRepo.FindByID(id)
}

// ListUsers returns all users.
func (s *UserService) ListUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

// UpdateUser validates and updates a user.
func (s *UserService) UpdateUser(user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	if user.ID == 0 {
		return errors.New("invalid user id")
	}

	if !utils.IsNonEmpty(user.Name) {
		return errors.New("user name is required")
	}

	if !utils.IsValidEmail(user.Email) {
		return errors.New("valid user email is required")
	}

	return s.userRepo.Update(user)
}

// DeleteUser removes a user by ID.
func (s *UserService) DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("invalid user id")
	}
	return s.userRepo.Delete(id)
}
