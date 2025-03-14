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
// Get user from context
func (r *VoteController) Store(ctx http.Context) http.Response {
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return ctx.Response().Json(http.StatusUnauthorized, models.ErrorResponse{
			Message: "Unauthorized",
			Errors:  "Invalid token",
		})
	}

	// Validate request
	var request requests.CreateVote
	if errors, err := ctx.Request().ValidateRequest(&request); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Validation error",
			Errors:  err.Error(),
		})
	} else if errors != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Validation error",
			Errors:  errors.All(),
		})
	}

	optionID, err := strconv.ParseUint(request.OptionID, 10, 64)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, models.ErrorResponse{
			Message: "Invalid option ID",
			Errors:  "Option ID must be a valid number",
		})
	}

	// Start transaction
	tx, err := facades.Orm().Query().Begin()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Internal server error",
			Errors:  "Failed to start database transaction",
		})
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get poll by code with a single query
	var poll models.Polls
	if err := tx.Where("code = ?", request.Code).First(&poll); err != nil {
		tx.Rollback()
		return ctx.Response().Json(http.StatusNotFound, models.ErrorResponse{
			Message: "Poll not found",
			Errors:  "The requested poll does not exist",
		})
	}

	// Check if poll is active
	if poll.Status != models.Active {
		tx.Rollback()
		return ctx.Response().Json(http.StatusConflict, models.ErrorResponse{
			Message: "Poll is not active",
			Errors:  "Cannot vote on an inactive poll",
		})
	}

	// Check if option exists and belongs to the poll in a single query
	var option models.Options
	if err := tx.Where("id = ? AND poll_id = ?", optionID, poll.ID).FirstOrFail(&option); err != nil {
		tx.Rollback()
		return ctx.Response().Json(http.StatusNotFound, models.ErrorResponse{
			Message: "Option not found",
			Errors:  "The selected option is invalid for this poll",
		})
	}

	// Check if user has already voted
	var hasVoted bool
	if err := tx.Model(&models.Votes{}).Where("user_id = ? AND poll_id = ?", user.ID, poll.ID).Exists(&hasVoted); err != nil {
		tx.Rollback()
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to check vote status",
			Errors:  "Database error occurred when checking vote status",
		})
	}

	if hasVoted {
		tx.Rollback()
		return ctx.Response().Json(http.StatusConflict, models.ErrorResponse{
			Message: "You have already voted in this poll",
			Errors:  "Each user may only vote once per poll",
		})
	}

	// Create vote record
	vote := models.Votes{
		UserID:   user.ID,
		PollID:   poll.ID,
		OptionID: uint(optionID),
	}

	if err := tx.Create(&vote); err != nil {
		tx.Rollback()
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to record vote",
			Errors:  "Database error occurred when saving your vote",
		})
	}

	// Update vote count with a direct SQL update for better concurrency
	_, err = tx.Model(&models.Options{}).
		Where("id = ?", optionID).
		Update("votes_count", option.VotesCount+1)
	if err != nil {
		tx.Rollback()
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to update vote count",
			Errors:  "Database error occurred when updating vote totals",
		})
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Failed to complete voting process",
			Errors:  "Database transaction could not be committed",
		})
	}

	return ctx.Response().Json(http.StatusCreated, models.ResponseWithMessage{
		Message: "Vote recorded successfully",
	})
}
