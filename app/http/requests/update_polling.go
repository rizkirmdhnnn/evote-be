package requests

import (
	"errors"
	"evote-be/app/models"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdatePolling struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	StartDate   time.Time     `json:"start_date"`
	EndDate     time.Time     `json:"end_date"`
	Status      models.Status `json:"status"`
}

func (r *UpdatePolling) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdatePolling) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdatePolling) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"title":       "string",
		"description": "string",
		"start_date":  "date",
		"end_date":    "date",
		"status":      "in:active,done",
	}
}

func (r *UpdatePolling) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdatePolling) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdatePolling) PrepareForValidation(ctx http.Context, data validation.Data) error {
	localNow := time.Now()

	// Check if start_date and end_date are provided
	startDateStr, exists := data.Get("start_date")
	if exists {
		StartDate, err := time.Parse(time.RFC3339, startDateStr.(string))
		if err != nil {
			return errors.New("invalid start_date format, use RFC3339 format")
		}
		r.StartDate = StartDate
	}

	endDateStr, exists := data.Get("end_date")
	if exists {
		endDate, err := time.Parse(time.RFC3339, endDateStr.(string))
		if err != nil {
			return errors.New("invalid end_date format, use RFC3339 format")
		}
		r.EndDate = endDate
	}

	// Check if start_date and end_date are valid
	if r.StartDate.After(r.EndDate) {
		return errors.New("tanggal selesai harus setelah tanggal mulai")
	}

	// Check if start_date is in the future
	if !r.StartDate.After(localNow) {
		return errors.New("tanggal mulai harus di masa depan")
	}

	return nil
}
