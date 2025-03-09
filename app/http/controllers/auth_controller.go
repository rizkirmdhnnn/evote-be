package controllers

import (
	"errors"
	"evote-be/app/http/requests"
	"evote-be/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/jackc/pgx/v5/pgconn"

	_ "evote-be/docs"
)

const UniqueViolation = "23505"

type AuthController struct {
	// Dependent services
}

func NewAuthController() *AuthController {
	return &AuthController{
		// Inject services
	}
}

// Register new user
//
// @Summary     Register new user
// @Description Register a new user account with a unique email address.
//
//	If the email is already taken or if the input is invalid,
//	appropriate error messages are returned.
//
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       request body requests.UserRegister true "User Registration Data"
// @Success 	 201 {object} models.ResponseWithData[models.UserRegisterResponse] "Success response"
// @Failure     400 {object} models.ErrorResponse "Validation error or email already taken"
// @Failure     500 {object} models.ErrorResponse "Internal server error"
// @Router      /auth/register [post]
func (r *AuthController) Register(ctx http.Context) http.Response {
	// Validate request data
	var req requests.UserRegister
	allerror, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "validation error",
			"errors":  err.Error(),
		})
	}
	if allerror != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "validation error",
			"errors":  allerror.All(),
		})
	}

	// Hash password
	hashedPass, err := facades.Hash().Make(req.Password)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  err.Error(),
		})
	}

	// Create user
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPass,
	}

	// Save user
	if err := facades.Orm().Query().Create(&user); err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			// Check if email already taken
			if pgError.Code == UniqueViolation && pgError.ConstraintName == "users_email_unique" {
				return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
					Message: "ups, something went wrong",
					Errors:  http.Json{"email": "email already taken"},
				})
			}
		}

		// Return error in case of any other error
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  err.Error(),
		})
	}

	// Return success response
	return ctx.Response().Json(http.StatusCreated, models.ResponseWithData[models.UserRegisterResponse]{
		Message: "user registered successfully",
		Data: models.UserRegisterResponse{
			ID:     int(user.ID),
			Name:   user.Name,
			Email:  user.Email,
			Avatar: user.Avatar,
		},
	})

}

// @Summary     Login user
//
// @Description Login user with email and password
//
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       request body requests.UserLogin true "User Login Data"
// @Success 	 200 {object} models.ResponseWithData[models.UserLoginResponse] "Success response"
// @Failure     400 {object} models.ErrorResponse "Validation error"
// @Failure    401 {object} models.ErrorResponse "Unauthorized"
// @Failure     500 {object} models.ErrorResponse "Internal server error"
// @Router      /auth/login [post]
func (r *AuthController) Login(ctx http.Context) http.Response {
	// Validate request data
	var req requests.UserLogin
	errors, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "validation error",
			"errors":  err.Error(),
		})
	}
	if errors != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "validation error",
			"errors":  errors.All(),
		})
	}

	// Find user by email
	user := models.User{}
	if err := facades.Orm().Query().Where("email", req.Email).First(&user); err != nil {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  http.Json{"email": "email not found"},
		})
	}

	// Check password
	if !facades.Hash().Check(req.Password, user.Password) {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  http.Json{"password": "password not match"},
		})
	}

	// Generate token
	token, err := facades.Auth(ctx).LoginUsingID(user.ID)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  err.Error(),
		})
	}

	// Return success response
	return ctx.Response().Json(http.StatusOK, models.ResponseWithData[models.UserLoginResponse]{
		Message: "user logged in successfully",
		Data: models.UserLoginResponse{
			ID:     int(user.ID),
			Name:   user.Name,
			Email:  user.Email,
			Avatar: user.Avatar,
			Token:  token,
		},
	})
}
