package seeders

import (
	"evote-be/app/models"

	"github.com/goravel/framework/facades"
)

type OptionSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *OptionSeeder) Signature() string {
	return "OptionSeeder"
}

// Run executes the seeder logic.
func (s *OptionSeeder) Run() error {
	var options []models.Options

	options = append(options, models.Options{
		PollID: 1,
		Name:   "Option 1",
		Desc:   "Option 1 Description",
	})

	options = append(options, models.Options{
		PollID: 1,
		Name:   "Option 2",
		Desc:   "Option 2 Description",
	})

	return facades.Orm().Query().Create(&options)
}
