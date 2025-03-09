package requests

import (
	"errors"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreatePolling struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date" swaggertype:"string" format:"date-time"`
	EndDate     time.Time `json:"end_date" swaggertype:"string" format:"date-time"`
	// testing:
	// * active - Active, can be voted
	// * done - Done, can't be voted
	Status string `json:"status" swaggertype:"string" enums:"active,done"`
}

func (r *CreatePolling) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreatePolling) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreatePolling) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"title":       "required|string",
		"description": "required|string",
		"start_date":  "required|date",
		"end_date":    "required|date",
		"status":      "in:active,done",
	}
}

func (r *CreatePolling) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreatePolling) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreatePolling) PrepareForValidation(ctx http.Context, data validation.Data) error {
	localNow := time.Now()

	// Check if start_date and end_date are provided
	startDateStr, exists := data.Get("start_date")
	if !exists {
		return errors.New("start_date is required")
	}
	endDateStr, exists := data.Get("end_date")
	if !exists {
		return errors.New("end_date is required")
	}

	// Parse start_date and end_date without changing timezone
	startDate, err := time.Parse(time.RFC3339, startDateStr.(string))
	if err != nil {
		return errors.New("invalid start_date format, use RFC3339 format")
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr.(string))
	if err != nil {
		return errors.New("invalid end_date format, use RFC3339 format")
	}

	// Store parsed times to struct
	r.StartDate = startDate
	r.EndDate = endDate

	// Check if start_date is after end_date
	if startDate.After(endDate) {
		return errors.New("tanggal selesai harus setelah tanggal mulai")
	}

	// Check if start_date is in the future
	if !startDate.After(localNow) {
		return errors.New("tanggal mulai harus di masa depan")
	}

	return nil
}
