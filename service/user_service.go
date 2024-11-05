package service

import (
	"errors"
	"fmt"

	"github.com/golang-web/model"
	"github.com/golang-web/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) RegisterNewUser(payload model.User) error {
	if payload.ID == "" || payload.Username == "" || payload.Password == "" || payload.Email == "" {
		return fmt.Errorf("all payload is required")
	}

	err := s.userRepo.Create(&payload)
	if err != nil {
		return fmt.Errorf("failed to create user: %s", err)
	}
	return nil
	// userRepo := repository.NewUserRepository(&payload)
}

func (cs *UserService) UserByID(id int) (*model.User, error) {

	user, err := cs.userRepo.UserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]*model.User, error) {
	users, err := s.userRepo.GetAllUser()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %v", err)
	}
	return users, nil
}

func (s *UserService) Login(username, password string) (*model.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password cannot be empty")
	}

	customer := model.User{
		Username: username,
		Password: password,
	}

	err := s.userRepo.Login(customer)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: customer.Username,
		Password: customer.Password,
		Email:    customer.Email,
	}

	return user, nil
}
