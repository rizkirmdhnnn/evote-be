package commands

import (
	"evote-be/app/models"
	"fmt"
	"time"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/facades"
)

type EndPoll struct {
}

// Signature The name and signature of the console command.
func (receiver *EndPoll) Signature() string {
	return "poll:end"
}

// Description The console command description.
func (receiver *EndPoll) Description() string {
	return "End polls that have reached their end date"
}

// Extend The console command extend.
func (receiver *EndPoll) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *EndPoll) Handle(ctx console.Context) error {
	now := time.Now().Truncate(time.Minute)
	layout := "2006-01-02 15:04"

	// Get all active polls that have passed their end date
	var polls []models.Polls
	err := facades.Orm().Query().
		Where("end_date <= ?", now.Format(layout)).
		Where("status = ?", "Active").
		Get(&polls)

	if err != nil {
		return err
	}

	// Update each poll's status to done
	for _, poll := range polls {
		result, err := facades.Orm().Query().Model(&poll).
			Where("id", poll.ID).
			Update("status", "Done")
		if err != nil {
			facades.Log().Error("Failed to end poll: " + err.Error())
		}
		if result.RowsAffected == 0 {
			facades.Log().Error("Failed to end poll: " + fmt.Sprint(poll.ID))
		}
	}

	return nil
}
