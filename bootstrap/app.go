package bootstrap

import (
	"github.com/goravel/framework/foundation"

	"evote-be/config"
)

func Boot() {
	app := foundation.NewApplication()

	// Bootstrap the application
	app.Boot()

	// Bootstrap the config.
	config.Boot()
}
