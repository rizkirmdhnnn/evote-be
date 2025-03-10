package controllers

import (
	"errors"
	"evote-be/app/http/requests"
	"evote-be/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserController struct {
	// Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		// Inject services
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

// Update Update user
// @Summary Update user
// @Description Update user
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body requests.UserUpdate true "User data"
// @Success 	200 {object} models.ResponseWithData[models.UserRegisterResponse] "Success response"
// @Failure 	400 {object} models.ErrorResponse "Validation error"
// @Failure 	401 {object} models.ErrorResponse "Unauthorized"
// @Router /users/update [put]
func (r *UserController) Update(ctx http.Context) http.Response {
	// Get user from context
	u, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// Validate request
	var request requests.UserUpdate
	if errors, err := ctx.Request().ValidateRequest(&request); err != nil || errors != nil {
		errorMsg := "Validation error"
		var errorData any = err
		if errors != nil {
			errorData = errors.All()
		} else if err != nil {
			errorData = err.Error()
		}

		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: errorMsg,
			Errors:  errorData,
		})
	}

	// Prepare user data
	user := models.User{
		Name:  u.Name,
		Email: u.Email,
	}
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}

	// Update user
	result, err := facades.Orm().Query().Model(&models.User{}).Where("id = ?", u.ID).Update(&user)
	if err != nil {
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
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Failed to update user",
			Errors:  err.Error(),
		})
	}
	// Check if user not found
	if result.RowsAffected == 0 {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Failed to update user",
			Errors:  "User not found",
		})
	}

	// Return success response
	return ctx.Response().Json(http.StatusCreated, models.ResponseWithData[models.UserRegisterResponse]{
		Message: "user updated successfully",
		Data: models.UserRegisterResponse{
			ID:     int(user.ID),
			Name:   user.Name,
			Email:  user.Email,
			Avatar: user.Avatar,
		},
	})
}
