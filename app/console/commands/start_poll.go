package commands

import (
	"evote-be/app/models"
	"fmt"
	"time"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/facades"
)

type StartPoll struct {
}

// Signature The name and signature of the console command.
func (receiver *StartPoll) Signature() string {
	return "poll:start"
}

// Description The console command description.
func (receiver *StartPoll) Description() string {
	return "Start polls that have reached their start date"
}

// Extend The console command extend.
func (receiver *StartPoll) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *StartPoll) Handle(ctx console.Context) error {
	now := time.Now().Truncate(time.Minute)
	layout := "2006-01-02 15:04"

	// Get all inactive polls that have passed their start date
	var polls []models.Polls
	err := facades.Orm().Query().Where("status", "Scheduled").
		Where("start_date >= ?", now.Format(layout)).
		Get(&polls)

	if err != nil {
		return err
	}

	// Update each poll's status to active
	for _, poll := range polls {
		result, err := facades.Orm().Query().Model(&poll).
			Where("id", poll.ID).
			Update("status", "Active")
		if err != nil {
			facades.Log().Error("Failed to start poll" + err.Error())
		}
		if result.RowsAffected == 0 {
			facades.Log().Error("Failed to start poll" + fmt.Sprint(poll.ID))
		}
	}
	return nil
}
