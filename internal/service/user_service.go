package service

import (
	"errors"

	"simple.market/internal/domain"
	"simple.market/internal/repository"
	"simple.market/pkg/utils"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(email, password, passwordConfirmation string) (*domain.User, error) {
	if !utils.IsValidEmail(email) {
		return nil, errors.New("invalid email")
	}

	existingUser, err := s.repo.FindByEmail(email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already in use")
	}

	if password != passwordConfirmation {
		return nil, errors.New("password and password confirmation doesnt match")
	}

	hashedPassword, err := utils.HashedPassword(password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &domain.User{
		Email:    email,
		Password: hashedPassword,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
