package database

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/contracts/database/seeder"

	"evote-be/database/migrations"
	"evote-be/database/seeders"
)

type Kernel struct {
}

func (kernel Kernel) Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20240915060148CreateUsersTable{},
		&migrations.M20250308113928CreatePollsTable{},
		&migrations.M20250308204957CreateOptionsTable{},
		&migrations.M20250308204808CreateVotesTable{},
	}
}

func (kernel Kernel) Seeders() []seeder.Seeder {
	return []seeder.Seeder{
		&seeders.DatabaseSeeder{},
		&seeders.UserSeeder{},
		&seeders.PollSeeder{},
		&seeders.OptionSeeder{},
	}
}
