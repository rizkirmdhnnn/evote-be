package routes

import (
	"evote-be/app/http/controllers"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	// Swagger
	swaggerController := controllers.NewSwaggerController()
	facades.Route().Get("/swagger/*any", swaggerController.Index)
}
