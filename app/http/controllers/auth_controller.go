package controllers

import (
	"errors"
	"evote-be/app/http/requests"
	"evote-be/app/mails"
	"evote-be/app/models"
	"fmt"
	"time"

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
		Name:              req.Name,
		Email:             req.Email,
		Password:          hashedPass,
		VerificationToken: randomString(32),
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

	// Send email
	err = facades.Mail().Queue(mails.NewUserRegister(user.Email, user.VerificationToken))
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  err.Error(),
		})
	}
	//

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
	var user models.User
	if err := facades.Orm().Query().Where("email = ?", req.Email).FirstOrFail(&user); err != nil {
		fmt.Print(err)
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

	// Check if user is verified
	if user.EmailVerifiedAt == "" {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "please verify your email address",
			Errors:  http.Json{"email": "email not verified"},
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

// @Summary     Verify email
// @Description Verify user email address and return an HTML page
// @Tags        Auth
// @Accept      json
// @Produce     text/html
// @Param       token path string true "Verification Token from email"
// @Success     200 {string} string "HTML content (success or error message)"
// @Router      /auth/verify/{token} [get]
func (r *AuthController) Verify(ctx http.Context) http.Response {
	// Get token from path
	token := ctx.Request().Route("token")

	// Find user by token
	var user models.User
	if err := facades.Orm().Query().Where("verification_token", token).FirstOrFail(&user); err != nil {
		return ctx.Response().View().Make("email-verify.tmpl", map[string]any{
			"success": false,
			"message": "The verification link has expired or is invalid. Please request a new one",
		})

	}

	// Check if user email already verified
	if user.EmailVerifiedAt != "" {
		return ctx.Response().View().Make("email-verify.tmpl", map[string]any{
			"status":  false,
			"message": "Your account has been already verified.",
		})

	}

	// Update user email verified at
	now := time.Now().Truncate(time.Minute)
	layout := "2006-01-02 15:04:05"
	user.EmailVerifiedAt = now.Format(layout)

	// Save user
	if err := facades.Orm().Query().Save(&user); err != nil {
		return ctx.Response().View().Make("email-verify.tmpl", map[string]any{
			"success": false,
			"message": "Oops! Something went wrong. Please try again.",
		})
	}

	return ctx.Response().View().Make("email-verify.tmpl", map[string]any{
		"success": true,
		"message": "Your account has been successfully verified. Welcome aboard! 🎉",
	})
}
