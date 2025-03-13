package controllers

import (
	"evote-be/app/http/requests"
	"evote-be/app/models"
	"strconv"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type VoteController struct {
}

func NewVoteController() *VoteController {
	return &VoteController{}
}

// @Summary Record a vote
// @Description Record a vote for a poll option
// @Tags Vote
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body requests.CreateVote true "Poll Data"
// @Router /votes/create [post]
func (r *VoteController) Store(ctx http.Context) http.Response {
	// Get user from context
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// Validate request
	var request requests.CreateVote
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

	// Parse poll_id and option_id
	pollID, err := strconv.ParseUint(request.PollID, 10, 64)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Invalid poll ID",
			Errors:  err.Error(),
		})
	}

	optionID, err := strconv.ParseUint(request.OptionID, 10, 64)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Invalid option ID",
			Errors:  err.Error(),
		})
	}

	// Check if user has already voted
	var alreadyVoted bool
	if err := facades.Orm().Query().Model(models.Votes{}).
		Where("user_id = ? AND poll_id = ?", user.ID, pollID).
		Exists(&alreadyVoted); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to check voting status",
			Errors:  err.Error(),
		})
	}
	if alreadyVoted {
		return ctx.Response().Json(http.StatusConflict, models.ErrorResponse{
			Message: "You have already voted in this poll",
			Errors:  "Duplicate vote",
		})
	}

	// Validate poll and option
	var count int64
	if err := facades.Orm().Query().
		Table("polls").
		Join("JOIN options ON options.poll_id = polls.id").
		Where("polls.id = ? AND polls.status = ? AND options.id = ?", pollID, models.Active, optionID).
		Count(&count); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to validate poll and option",
			Errors:  err.Error(),
		})
	}

	if count == 0 {
		return ctx.Response().Json(http.StatusNotFound, models.ErrorResponse{
			Message: "Poll not found, not active, or option doesn't belong to this poll",
			Errors:  "Invalid poll or option",
		})
	}

	// Get option for updating vote count
	var option models.Options
	if err := facades.Orm().Query().Where("id = ?", optionID).First(&option); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to retrieve option",
			Errors:  err.Error(),
		})
	}

	// Create vote
	vote := models.Votes{
		UserID:   user.ID,
		PollID:   uint(pollID),
		OptionID: uint(optionID),
	}

	// Start transaction
	tx, err := facades.Orm().Query().Begin()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to start transaction",
			Errors:  err.Error(),
		})
	}

	// Create vote
	if err := tx.Create(&vote); err != nil {
		tx.Rollback()
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to create vote",
			Errors:  err.Error(),
		})
	}

	// Increment votes_count in options table
	option.VotesCount++

	// Save option
	if err := tx.Save(&option); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to update vote count",
			Errors:  err.Error(),
		})
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to commit transaction",
			Errors:  err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusCreated, models.ResponseWithMessage{
		Message: "Vote recorded successfully",
	})
}
