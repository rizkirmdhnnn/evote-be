package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250308204808CreateVotesTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250308204808CreateVotesTable) Signature() string {
	return "20250308204808_create_votes_table"
}

// Up Run the migrations.
func (r *M20250308204808CreateVotesTable) Up() error {
	if !facades.Schema().HasTable("votes") {
		return facades.Schema().Create("votes", func(table schema.Blueprint) {
			table.BigIncrements("id")
			table.UnsignedBigInteger("user_id")
			table.UnsignedBigInteger("poll_id")
			table.UnsignedBigInteger("option_id")

			table.Foreign("user_id").References("id").On("users")
			table.Foreign("poll_id").References("id").On("polls")
			table.Foreign("option_id").References("id").On("options")
			table.Timestamps()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250308204808CreateVotesTable) Down() error {
	return facades.Schema().DropIfExists("votes")
}
