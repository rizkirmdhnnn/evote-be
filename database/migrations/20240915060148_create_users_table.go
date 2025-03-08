package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20240915060148CreateUsersTable struct {
}

// Signature The unique signature for the migration.
func (r *M20240915060148CreateUsersTable) Signature() string {
	return "20240915060148_create_users_table"
}

// Up Run the migrations.
func (r *M20240915060148CreateUsersTable) Up() error {
	return facades.Schema().Create("users", func(table schema.Blueprint) {
		table.BigIncrements("id")
		table.String("name")
		table.String("email")
		table.String("password")
		table.String("avatar").Default("")
		table.Timestamps()
		table.SoftDeletes()

		table.Unique("email")
	})
}

// Down Reverse the migrations.
func (r *M20240915060148CreateUsersTable) Down() error {
	return facades.Schema().DropIfExists("")
}
