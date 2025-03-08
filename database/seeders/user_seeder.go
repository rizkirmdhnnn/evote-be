package seeders

import (
	"evote-be/app/models"

	"github.com/goravel/framework/facades"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	user := models.User{
		Name:     "Rizkirmdhn",
		Email:    "achmadrizkiramadhan0101@gmail.com",
		Password: "password",
	}

	// Hash the password
	user.Password, _ = facades.Hash().Make(user.Password)

	// Create the user
	return facades.Orm().Query().Create(&user)
}
