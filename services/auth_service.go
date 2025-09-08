package services

import (
	"kasirku/models/domain"
)

type AuthService interface {
	Login(request domain.LoginRequest) (string, error)
}
