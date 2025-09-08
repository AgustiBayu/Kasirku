package helpers

import (
	"kasirku/models/domain"
	"kasirku/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)

	// Check if superadmin already exists
	_, err := userRepo.FindByUsername("superadmin")
	if err == nil {
		// User already exists
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	users := []domain.User{
		{Name: "Super Admin", Username: "superadmin", Password: string(hashedPassword), Role: "superadmin", IsActive: true},
		{Name: "Admin User", Username: "admin", Password: string(hashedPassword), Role: "admin", IsActive: true},
		{Name: "Cashier User", Username: "cashier", Password: string(hashedPassword), Role: "cashier", IsActive: true},
	}

	for _, user := range users {
		db.Create(&user)
	}
}
