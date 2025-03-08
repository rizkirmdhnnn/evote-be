package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250308113928CreatePollsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250308113928CreatePollsTable) Signature() string {
	return "20250308113928_create_polls_table"
}

// Up Run the migrations.
func (r *M20250308113928CreatePollsTable) Up() error {
	if !facades.Schema().HasTable("polls") {
		return facades.Schema().Create("polls", func(table schema.Blueprint) {
			table.BigIncrements("id")
			table.String("title")
			table.String("description")
			table.String("status")
			table.Timestamp("start_date")
			table.Timestamp("end_date")
			table.Timestamps()
			table.SoftDeletes()
			table.BigInteger("user_id").Unsigned()

			table.Foreign("user_id").References("id").On("users")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250308113928CreatePollsTable) Down() error {
	return facades.Schema().DropIfExists("polls")
}
