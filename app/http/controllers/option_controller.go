package controllers

import (
	"evote-be/app/http/requests"
	"evote-be/app/models"
	"fmt"
	"strconv"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type OptionController struct {
	// Dependent services
}

func NewOptionController() *OptionController {
	return &OptionController{
		// Inject services
	}
}

func (r *OptionController) Index(ctx http.Context) http.Response {
	return nil
}

// Store Create a new option
//
// @Summary Create a new option
// @Description Create a new option
// @Tags Options
// @Accept json
// @Produce json
// @Security Bearer
// @Param request formData requests.CreateOption true "Option data"
// @Param avatar formData file false "Option avatar"
// @Success 201 {object} models.ResponseWithData[models.CreateOptionsResponse] "Option created"
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Router /options/create [post]
func (r *OptionController) Store(ctx http.Context) http.Response {
	// Get user from context
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// Validate request
	var request requests.CreateOption
	errors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Validation error",
			Errors:  err.Error(),
		})
	}
	if errors != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Validation error",
			Errors:  errors.All(),
		})
	}

	// Check if poll_id is valid
	pollID, err := strconv.ParseUint(request.PollID, 10, 64)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Validation error",
			Errors:  "Invalid poll_id",
		})
	}

	// Check if poll exists
	var poll models.Polls
	if err := facades.Orm().Query().Model(&poll).Where("id = ?", pollID).FirstOrFail(&poll); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "upss, something went wrong",
			Errors:  "poll not found",
		})
	}

	// Check if user is the owner of the poll
	if poll.UserID != user.ID {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "You are not the owner of this poll",
		})
	}

	// Get file
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

	// Generate file name
	fileName := fmt.Sprintf("poll_%d_option_%d.%s", poll.ID, user.ID, extension)

	// Upload file to MinIO
	path, err := facades.Storage().Disk("minio").PutFileAs("options", file, fileName)
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
	// Create new option
	option := models.Options{
		Name:   request.Name,
		Desc:   request.Desc,
		Avatar: url,
		PollID: uint(pollID),
	}

	// Save option
	if err := facades.Orm().Query().Create(&option); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "upss, something went wrong",
			Errors:  err.Error(),
		})
	}

	// Return response
	return ctx.Response().Json(http.StatusCreated, models.ResponseWithData[models.CreateOptionsResponse]{
		Message: "Option created",
		Data:    option.ToResponse(),
	})

}

// Update Update an option
// @Summary Update an option
// @Description Update an option
// @Tags Options
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Option ID"
// @Param request formData requests.UpdateOption true "Option data"
// @Param avatar formData file false "Option avatar"
// @Success 200 {object} models.ResponseWithData[models.CreateOptionsResponse] "Option updated"
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Option not found"
// @Router /options/{id}/update [put]
func (r *OptionController) Update(ctx http.Context) http.Response {
	// Get user from context
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// Get option id
	optionID := ctx.Request().Route("id")

	// Validate request
	var request requests.UpdateOption
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

	// Get avatar
	file, _ := ctx.Request().File("avatar")

	// Check if option_id is valid
	id, err := strconv.ParseUint(optionID, 10, 64)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Validation error",
			Errors:  "Invalid option_id",
		})
	}

	// Check if option exists
	var option models.Options
	query := facades.Orm().Query()
	if err := query.Model(&option).Where("id = ?", id).FirstOrFail(&option); err != nil {
		return ctx.Response().Json(http.StatusNotFound, models.ErrorResponse{
			Message: "Option not found",
			Errors:  "Option not found",
		})
	}

	// Check if poll exists
	var poll models.Polls
	if err := query.Model(&poll).Where("id = ?", option.PollID).FirstOrFail(&poll); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "upss, something went wrong",
			Errors:  "poll not found",
		})
	}

	// Check if user is the owner of the poll
	if poll.UserID != user.ID {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "You are not the owner of this poll",
		})
	}

	// Update option only if values are not empty
	if request.Name != "" {
		option.Name = request.Name
	}
	if request.Desc != "" {
		option.Desc = request.Desc
	}
	if file != nil {
		// Get file extension
		extension, err := file.Extension()
		if err != nil {
			return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
				Message: "Failed to determine file extension",
				Errors:  err.Error(),
			})
		}

		// Generate file name
		fileName := fmt.Sprintf("poll_%d_option_%d.%s", poll.ID, user.ID, extension)

		// Upload file to MinIO
		path, err := facades.Storage().Disk("minio").PutFileAs("options", file, fileName)
		if err != nil {
			return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
				Message: "Failed to upload avatar",
				Errors:  err.Error(),
			})
		}

		// Get URL from MinIO
		url := facades.Storage().Disk("minio").Url(path)

		// Fix URL schema
		if facades.Config().GetBool("MINIO_SSL") {
			url = "https://" + url
		} else {
			url = "http://" + url
		}

		// Update avatar
		option.Avatar = url
	}

	// Check if poll_id is valid
	pollID, err := strconv.ParseUint(request.PollID, 10, 64)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Validation error",
			Errors:  "Invalid poll_id",
		})
	}
	option.PollID = uint(pollID)

	// Save option
	if err := query.Save(&option); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to update option",
			Errors:  err.Error(),
		})
	}

	// Return response
	return ctx.Response().Json(http.StatusOK, models.ResponseWithData[models.CreateOptionsResponse]{
		Message: "Option updated successfully",
		Data:    option.ToResponse(),
	})
}

// Delete Delete an option
// @Summary Delete an option
// @Description Delete an option
// @Tags Options
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Option ID"
// @Success 200 {object} models.ResponseWithMessage "Option deleted"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Option not found"
// @Router /options/{id}/delete [delete]
func (r *OptionController) Delete(ctx http.Context) http.Response {
	// Get user from context
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// Get option id
	optionID := ctx.Request().Route("id")

	// Delete option
	result, err := facades.Orm().Query().
		Model(&models.Options{}).
		Where("options.id = ? AND EXISTS (SELECT 1 FROM polls WHERE polls.id = options.poll_id AND polls.user_id = ?)",
			optionID, user.ID).
		Delete()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to delete option",
			Errors:  err.Error(),
		})
	}

	// Check if any row was affected (option existed and user owned it)
	if result.RowsAffected == 0 {
		return ctx.Response().Json(http.StatusNotFound, models.ErrorResponse{
			Message: "Option not found",
			Errors:  "Option not found or you don't have permission to delete it",
		})
	}

	// Return response
	return ctx.Response().Json(http.StatusOK, models.ResponseWithMessage{
		Message: "Option deleted successfully",
	})
}
