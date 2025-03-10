package seeders

import (
	"evote-be/app/models"
	"time"

	"github.com/goravel/framework/facades"
)

type PollSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *PollSeeder) Signature() string {
	return "PollSeeder"
}

// Run executes the seeder logic.
func (s *PollSeeder) Run() error {
	polls := models.Polls{
		Title:       "Polling 1",
		Description: "Polling 1 Description",
		Status:      models.Active,
		StartDate:   time.Now().AddDate(0, 0, 1),
		EndDate:     time.Now().AddDate(0, 0, 2),
		UserID:      1,
	}
	return facades.Orm().Query().Create(&polls)
}
