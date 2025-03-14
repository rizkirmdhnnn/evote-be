package console

import (
	"evote-be/app/console/commands"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/schedule"
	"github.com/goravel/framework/facades"
)

type Kernel struct {
}

func (kernel *Kernel) Commands() []console.Command {
	return []console.Command{
		&commands.EndPoll{},
		&commands.StartPoll{},
	}
}

func (kernel *Kernel) Schedule() []schedule.Event {
	return []schedule.Event{
		facades.Schedule().Command("poll:end").EveryMinute(),
		facades.Schedule().Command("poll:start").EveryMinute(),
	}
}
