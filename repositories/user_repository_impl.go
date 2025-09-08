package repositories

import (
	"kasirku/models/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) Save(user domain.User) (domain.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}
