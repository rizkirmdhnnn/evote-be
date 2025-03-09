package routes

import (
	"github.com/goravel/framework/facades"

	"evote-be/app/http/controllers"
	"evote-be/app/http/middleware"
)

func Api() {
	authController := controllers.NewAuthController()
	pollsController := controllers.NewPollsController()
	optionController := controllers.NewOptionController()

	// @Group Auth
	facades.Route().Post("/auth/register", authController.Register)
	facades.Route().Post("/auth/login", authController.Login)

	// @Group Polls
	facades.Route().Middleware(middleware.Auth()).Get("/polls", pollsController.Index)
	facades.Route().Middleware(middleware.Auth()).Post("/polls/create", pollsController.Store)
	facades.Route().Middleware(middleware.Auth()).Get("/polls/{id}", pollsController.Show)
	facades.Route().Middleware(middleware.Auth()).Put("/polls/{id}/update", pollsController.Update)
	facades.Route().Middleware(middleware.Auth()).Delete("/polls/{id}/delete", pollsController.Delete)

	// @Group Options
	facades.Route().Middleware(middleware.Auth()).Post("/options/create", optionController.Store)
	facades.Route().Middleware(middleware.Auth()).Delete("/options/{id}/delete", optionController.Delete)
}
