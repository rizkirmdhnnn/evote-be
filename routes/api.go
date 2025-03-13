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
	userController := controllers.NewUserController()
	voteController := controllers.NewVoteController()

	// @Group Auth
	facades.Route().Post("/auth/register", authController.Register)
	facades.Route().Post("/auth/login", authController.Login)

	// @Group Users
	facades.Route().Middleware(middleware.Auth()).Put("/users/update", userController.Update)
	facades.Route().Middleware(middleware.Auth()).Post("/users/avatar", userController.UploadAvatar)
	facades.Route().Middleware(middleware.Auth()).Get("/users/profile", userController.GetProfile)

	// @Group Polls
	facades.Route().Middleware(middleware.Auth()).Get("/polls", pollsController.Index)
	facades.Route().Middleware(middleware.Auth()).Post("/polls/create", pollsController.Store)
	facades.Route().Middleware(middleware.Auth()).Get("/polls/{id}", pollsController.Show)
	facades.Route().Middleware(middleware.Auth()).Put("/polls/{id}/update", pollsController.Update)
	facades.Route().Middleware(middleware.Auth()).Delete("/polls/{id}/delete", pollsController.Delete)
	facades.Route().Middleware(middleware.Auth()).Get("/polls/{id}/options", pollsController.GetPollOptions)
	facades.Route().Middleware(middleware.Auth()).Get("/polls/{id}/generate", pollsController.GeneratePublicPollCode)
	facades.Route().Get("/polls/public", pollsController.GetPublicPolls)

	// @Group Options
	facades.Route().Middleware(middleware.Auth()).Post("/options/create", optionController.Store)
	facades.Route().Middleware(middleware.Auth()).Delete("/options/{id}/delete", optionController.Delete)
	facades.Route().Middleware(middleware.Auth()).Put("/options/{id}/update", optionController.Update)

	// @Group Votes
	facades.Route().Middleware(middleware.Auth()).Post("/votes/create", voteController.Store)
}
