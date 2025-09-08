package services

import (
	"errors"
	"kasirku/helpers"
	"kasirku/models/domain"
	"kasirku/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &AuthServiceImpl{userRepo: userRepo}
}

func (s *AuthServiceImpl) Login(request domain.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByUsername(request.Username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if !user.IsActive {
		return "", errors.New("user is not active")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := helpers.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
