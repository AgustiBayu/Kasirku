package repositories

import (
	"kasirku/models/domain"
)

type UserRepository interface {
	FindByUsername(username string) (domain.User, error)
	Save(user domain.User) (domain.User, error)
}
