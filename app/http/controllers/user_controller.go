package controllers

import (
	"errors"
	"evote-be/app/http/requests"
	"evote-be/app/models"
	"fmt"

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

	// Get user data
	var user models.User
	if err := facades.Orm().Query().Model(&models.User{}).Where("id = ?", u.ID).First(&user); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Failed to update user",
			Errors:  "User not found",
		})
	}

	// Update fields if not empty
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

// UploadAvatar Upload user avatar
// @Summary Upload user avatar
// @Description Upload user avatar
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Param avatar formData file true "User avatar"
// @Success 	200 {object} models.ResponseWithData[models.UserRegisterResponse] "Success response"
// @Failure 	400 {object} models.ErrorResponse "Validation error"
// @Failure 	401 {object} models.ErrorResponse "Unauthorized"
// @Router /users/avatar [post]
// TODO: Compress image before storing
func (r *UserController) UploadAvatar(ctx http.Context) http.Response {
	// Get user from context
	u, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	file, err := ctx.Request().File("avatar")
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Failed to upload avatar",
			Errors:  err.Error(),
		})
	}

	// Get file extension
	extension, err := file.Extension()
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Failed to determine file extension",
			Errors:  err.Error(),
		})
	}

	// Allowed file types
	allowedTypes := map[string]bool{"jpg": true, "jpeg": true, "png": true, "gif": true}

	// Check if extension is allowed
	if !allowedTypes[extension] {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Invalid file type",
			Errors:  "Only JPG, JPEG, PNG, and GIF are allowed",
		})
	}

	// Generate file name
	fileName := fmt.Sprintf("avatar_%d.%s", u.ID, extension)

	// Upload file to MinIO
	path, err := facades.Storage().Disk("minio").PutFileAs("avatars", file, fileName)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to upload avatar",
			Errors:  err.Error(),
		})
	}

	// Ambil URL dari MinIO
	url := facades.Storage().Disk("minio").Url(path)

	// Perbaiki skema URL
	if facades.Config().GetBool("MINIO_SSL") {
		url = "https://" + url
	} else {
		url = "http://" + url
	}

	// Update user avatar
	result, err := facades.Orm().Query().Model(&models.User{}).Where("id = ?", u.ID).Update("avatar", url)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Failed to update user avatar",
			Errors:  err.Error(),
		})
	}

	// Check if user not found
	if result.RowsAffected == 0 {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Failed to update user avatar",
			Errors:  "User not found",
		})
	}

	// Get user data
	if err := facades.Orm().Query().Model(&models.User{}).Where("id = ?", u.ID).First(&u); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Failed to update user avatar",
			Errors:  "User not found",
		})
	}

	// Return success response
	return ctx.Response().Json(http.StatusCreated, models.ResponseWithData[models.UserRegisterResponse]{
		Message: "avatar uploaded successfully",
		Data: models.UserRegisterResponse{
			ID:     int(u.ID),
			Name:   u.Name,
			Email:  u.Email,
			Avatar: url,
		},
	})
}
