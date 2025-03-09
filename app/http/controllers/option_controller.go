package controllers

import (
	"evote-be/app/http/requests"
	"evote-be/app/models"
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
// @Param request body requests.CreateOption true "Option data"
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

	// Create new option
	option := models.Options{
		Name:   request.Name,
		Desc:   request.Desc,
		Avatar: request.Avatar,
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
// @Param request body requests.UpdateOption true "Option data"
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
	if request.Avatar != "" {
		option.Avatar = request.Avatar
	}
	if request.PollID != "" {
		pollID, err := strconv.ParseUint(request.PollID, 10, 64)
		if err != nil {
			return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
				Message: "Validation error",
				Errors:  "Invalid poll_id",
			})
		}
		option.PollID = uint(pollID)
	}

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
