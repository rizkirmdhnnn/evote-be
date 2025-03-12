package controllers

import (
	"evote-be/app/http/requests"
	"evote-be/app/models"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type PollsController struct {
	// Dependent services
}

func NewPollsController() *PollsController {
	return &PollsController{
		// Inject services
	}
}

// Get all polls
//
// @Summary     Get all polls
// @Description Get all polls
// @Tags        Polls
// @Accept      json
// @Produce     json
// @Security  Bearer
// @Param       limit query int false "Limit"
// @Param       offset query int false "Offset"
// @Success 	 200 {object} models.ResponseWithData[[]models.PollsResponse] "Success response"
// @Failure    	401 {object} models.ErrorResponse "Unauthorized"
// @Failure     500 {object} models.ErrorResponse "Internal server error"
// @Router      /polls [get]
func (r *PollsController) Index(ctx http.Context) http.Response {
	// Get user from context
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// Get query params
	limit := ctx.Request().QueryInt("limit", 10)
	offset := ctx.Request().QueryInt("offset", 0)

	// Get polls with optimized query
	var polls []models.Polls
	query := facades.Orm().Query().Model(&models.Polls{}).Where("user_id", user.ID).OrderBy("id", "desc")
	if err := query.Limit(limit).Offset(offset).Select("id", "title", "description", "status", "start_date", "end_date").Find(&polls); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Oops, something went wrong",
			Errors:  err.Error(),
		})
	}

	// Get total count
	var total int64
	if err := query.Count(&total); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Oops, something went wrong",
			Errors:  err.Error(),
		})
	}

	// Convert to response
	pollsResp := make([]models.PollsResponse, len(polls))
	for i, poll := range polls {
		pollsResp[i] = poll.ToResponse()
	}

	// Return response with optimized pagination
	return ctx.Response().Json(http.StatusOK, models.PaginateResponse[[]models.PollsResponse]{
		Message: "Polls fetched successfully",
		Data:    pollsResp,
		Meta: models.Meta{
			Total:    int(total),
			PerPage:  limit,
			LastPage: int(math.Ceil(float64(total) / float64(limit))),
			CurrPage: (offset / limit) + 1,
		},
	})
}

// Store new poll
//
// @Summary     Store new poll
// @Description Create new poll
// @Tags        Polls
// @Accept      json
// @Produce     json
// @Security  Bearer
// @Param       request body requests.CreatePolling true "Poll Data"
// @Success 	 201 {object} models.ResponseWithData[models.CreatePollingResponse] "Success response"
// @Failure    	401 {object} models.ErrorResponse "Unauthorized"
// @Failure     400 {object} models.ErrorResponse "Validation error or title already taken"
// @Failure     500 {object} models.ErrorResponse "Internal server error"
// @Router      /polls/create [post]
func (r *PollsController) Store(ctx http.Context) http.Response {
	// get user from context
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// validate request
	var request requests.CreatePolling
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

	// create poll object
	poll := models.Polls{
		Title:       request.Title,
		Description: request.Description,
		Status:      models.Active,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		UserID:      user.ID,
	}

	// create poll
	if err := facades.Orm().Query().Create(&poll); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  err.Error(),
		})
	}

	// return response
	return ctx.Response().Json(http.StatusCreated, models.ResponseWithData[models.CreatePollingResponse]{
		Message: "Poll created successfully",
		Data: models.CreatePollingResponse{
			ID:          int(poll.ID),
			Title:       poll.Title,
			Description: poll.Description,
			Status:      models.Active,
			StartDate:   poll.StartDate,
			EndDate:     poll.EndDate,
		},
	})
}

// Show poll
// @Summary     Show poll
// @Description Show poll
// @Tags        Polls
// @Accept      json
// @Produce     json
// @Security  Bearer
// @Param       id path int true "Poll ID"
// @Success 	 200 {object} models.ResponseWithData[models.PollsResponse] "Success response"
// @Failure    	401 {object} models.ErrorResponse "Unauthorized"
// @Failure     404 {object} models.ErrorResponse "Poll not found"
// @Failure     500 {object} models.ErrorResponse "Internal server error"
// @Router      /polls/{id} [get]
func (r *PollsController) Show(ctx http.Context) http.Response {
	// get user from context
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// get poll id from path
	id := ctx.Request().Route("id")

	// get poll
	var poll models.Polls
	if err := facades.Orm().Query().Where("user_id = ? AND id = ?", user.ID, id).FirstOrFail(&poll); err != nil {
		return ctx.Response().Json(http.StatusNotFound, models.ErrorResponse{
			Message: "Poll not found",
			Errors:  err.Error(),
		})
	}

	// return response
	return ctx.Response().Json(http.StatusOK, models.ResponseWithData[models.PollsResponse]{
		Message: "Poll fetched successfully",
		Data:    poll.ToResponse(),
	})
}

// Update poll
// @Summary     Update poll
// @Description Update poll
// @Tags        Polls
// @Accept      json
// @Produce     json
// @Security  Bearer
// @Param       id path int true "Poll ID"
// @Param       request body requests.UpdatePolling true "Poll Data"
// @Success 	 200 {object} models.ResponseWithData[models.UpdatePollingResponse] "Success response"
// @Failure    	401 {object} models.ErrorResponse "Unauthorized"
// @Failure     400 {object} models.ErrorResponse "Validation error or title already taken"
// @Failure     404 {object} models.ErrorResponse "Poll not found"
// @Failure     500 {object} models.ErrorResponse "Internal server error"
// @Router      /polls/{id}/update [put]
func (r *PollsController) Update(ctx http.Context) http.Response {
	// get user from context
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// get poll id from path
	id := ctx.Request().Route("id")

	// get poll
	var poll models.Polls
	if err := facades.Orm().Query().Where("user_id = ? AND id = ?", user.ID, id).FirstOrFail(&poll); err != nil {
		return ctx.Response().Json(http.StatusNotFound, models.ErrorResponse{
			Message: "Poll not found",
			Errors:  err.Error(),
		})
	}

	// validate request
	var request requests.UpdatePolling
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

	// update poll if value changed
	if request.Title != "" {
		poll.Title = request.Title
	}
	if request.Description != "" {
		poll.Description = request.Description
	}
	if request.StartDate.After(poll.StartDate) {
		poll.StartDate = request.StartDate
	}
	if request.EndDate.After(poll.EndDate) {
		poll.EndDate = request.EndDate
	}
	if request.Status != "" {
		poll.Status = request.Status
	}

	// save poll
	if err := facades.Orm().Query().Save(&poll); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  err.Error(),
		})
	}

	// return response
	return ctx.Response().Json(http.StatusOK, models.ResponseWithData[models.UpdatePollingResponse]{
		Message: "Poll updated successfully",
		Data: models.UpdatePollingResponse{
			ID:          int(poll.ID),
			Title:       poll.Title,
			Description: poll.Description,
			StartDate:   poll.StartDate,
			Status:      poll.Status,
			EndDate:     poll.EndDate,
		},
	})
}

// Delete poll
// @Summary     Delete poll
// @Description Delete poll
// @Tags        Polls
// @Accept      json
// @Produce     json
// @Security  Bearer
// @Param       id path int true "Poll ID"
// @Success 	 200 {object} models.ResponseWithMessage "Success response"
// @Failure    	401 {object} models.ErrorResponse "Unauthorized"
// @Failure     404 {object} models.ErrorResponse "Poll not found"
// @Failure     500 {object} models.ErrorResponse "Internal server error"
// @Router      /polls/{id}/delete [delete]
func (r *PollsController) Delete(ctx http.Context) http.Response {
	// get user from context
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}
	// get poll id from path
	id := ctx.Request().Route("id")

	// Check if poll exists and belongs to user
	var poll models.Polls
	if err := facades.Orm().Query().Model(&poll).Where("id = ? AND user_id = ?", id, user.ID).FirstOrFail(&poll); err != nil {
		return ctx.Response().Json(http.StatusNotFound, models.ErrorResponse{
			Message: "something went wrong",
			Errors:  "poll not found or you don't have permission",
		})
	}

	// Begin transaction
	tx, err := facades.Orm().Query().Begin()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  err.Error(),
		})
	}

	// Soft delete options
	if _, err := tx.Model(&models.Options{}).Where("poll_id = ?", id).Delete(); err != nil {
		tx.Rollback()
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  err.Error(),
		})
	}

	// Soft delete poll
	if _, err := tx.Model(&poll).Delete(); err != nil {
		tx.Rollback()
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  err.Error(),
		})
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "ups, something went wrong",
			Errors:  err.Error(),
		})
	}

	// return response
	return ctx.Response().Json(http.StatusOK, models.ResponseWithMessage{
		Message: "Poll deleted successfully",
	})
}

// Get all options of a poll
// @Summary Get all options of a poll
// @Description Get all options of a poll
// @Tags Polls
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Poll ID"
// @Success 200 {object} models.ResponseWithData[models.CreateOptionsResponse] "Options found"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Poll not found"
// @Router /polls/{id}/options [get]
func (r *PollsController) GetPollOptions(ctx http.Context) http.Response {
	// Get user from context
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// Get poll id
	pollID := ctx.Request().Route("id")

	// Check if poll_id is valid
	id, err := strconv.ParseUint(pollID, 10, 64)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Validation error",
			Errors:  "Invalid poll_id",
		})
	}

	// Check if poll exists
	var poll models.Polls
	if err := facades.Orm().Query().Model(&poll).Where("id = ?", id).FirstOrFail(&poll); err != nil {
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

	// Get all options of the poll
	var options []models.Options
	if err := facades.Orm().Query().Model(&options).Where("poll_id = ?", id).Scan(&options); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to get options",
			Errors:  err.Error(),
		})
	}

	// Convert options to response
	optResp := make([]models.CreateOptionsResponse, len(options))
	for i, opt := range options {
		optResp[i] = opt.ToResponse()
	}

	// Return response
	return ctx.Response().Json(http.StatusOK, models.ResponseWithData[[]models.CreateOptionsResponse]{
		Message: "Options found",
		Data:    optResp,
	})
}

// Get public polls, options for voting
// @Summary Get public polls, options for voting
//
// @Description Get public polls, options for voting
// @Tags Polls
// @Accept json
// @Produce json
// @Param code query string true "Poll Code"
// @Success 200 {object} models.ResponseWithData[models.PublicPollsResponse] "Polls found"
// @Failure 404 {object} models.ErrorResponse "Poll not found"
// @Router /polls/public [get]
func (r *PollsController) GetPublicPolls(ctx http.Context) http.Response {
	// Get query params from request
	code := ctx.Request().Query("code")

	// Check if code is empty
	if code == "" {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Validation error",
			Errors:  "Code is required",
		})
	}

	// Get poll and options by code
	var poll models.Polls
	if err := facades.Orm().Query().Model(&models.Polls{}).With("Options").Where("code = ?", code).FirstOrFail(&poll); err != nil {
		return ctx.Response().Json(http.StatusNotFound, models.ErrorResponse{
			Message: "Poll not found",
			Errors:  err.Error(),
		})
	}

	// Check if poll is active
	if poll.Status != models.Active {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "The poll is currently inactive.",
			Errors:  "POLL_INACTIVE",
		})
	}

	// Convert poll to Public Polls Response
	pollResp := poll.ToPublicResponse()

	// Return response
	return ctx.Response().Json(http.StatusOK, models.ResponseWithData[models.PublicPollsResponse]{
		Message: "Polls found",
		Data:    pollResp,
	})
}

// Generate public poll code
// @Summary Generate public poll code
// @Description Generate public poll code
// @Tags Polls
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Poll ID"
// @Success 200 {object} models.ResponseWithData[models.PollsResponse] "Poll code generated"
// @Failure 404 {object} models.ErrorResponse "Poll not found"
// @Router /polls/{id}/generate [get]
func (r *PollsController) GeneratePublicPollCode(ctx http.Context) http.Response {
	// Get user from context
	user, ok := ctx.Value("user").(models.User)
	fmt.Sprintf("user: %v", user)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// Get poll id from path
	id := ctx.Request().Route("id")

	// Check if poll exists and belongs to user
	var poll models.Polls
	if err := facades.Orm().Query().Model(&poll).Where("id = ? AND user_id = ?", id, user.ID).FirstOrFail(&poll); err != nil {
		return ctx.Response().Json(http.StatusNotFound, models.ErrorResponse{
			Message: "Poll not found",
			Errors:  err.Error(),
		})
	}

	// Generate public code
	if poll.Code != "" {
		return ctx.Response().Json(http.StatusOK, models.ResponseWithData[models.PollsResponse]{
			Message: "Poll code already generated",
			Data:    poll.ToResponse(),
		})
	}
	code := randomString(6)

	// Update poll with code
	poll.Code = code
	if err := facades.Orm().Query().Save(&poll); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to generate code",
			Errors:  err.Error(),
		})
	}

	// Return response
	return ctx.Response().Json(http.StatusOK, models.ResponseWithData[models.PollsResponse]{
		Message: "Poll code generated",
		Data:    poll.ToResponse(),
	})
}

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}
