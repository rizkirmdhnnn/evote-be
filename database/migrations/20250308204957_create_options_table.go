package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250308204957CreateOptionsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250308204957CreateOptionsTable) Signature() string {
	return "20250308204957_create_options_table"
}

// Up Run the migrations.
func (r *M20250308204957CreateOptionsTable) Up() error {
	if !facades.Schema().HasTable("options") {
		return facades.Schema().Create("options", func(table schema.Blueprint) {
			table.BigIncrements("id")
			table.String("name")
			table.String("desc")
			table.String("avatar").Nullable()
			table.UnsignedBigInteger("poll_id")
			table.UnsignedBigInteger("votes_count").Default(0)
			table.Timestamps()
			table.SoftDeletes()
			table.Foreign("poll_id").References("id").On("polls").CascadeOnDelete()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250308204957CreateOptionsTable) Down() error {
	return facades.Schema().DropIfExists("options")
}
