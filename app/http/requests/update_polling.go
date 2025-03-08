package requests

import (
	"errors"
	"evote-be/app/models"
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdatePolling struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	StartDate   carbon.DateTime `json:"start_date"`
	EndDate     carbon.DateTime `json:"end_date"`
	Status      models.Status   `json:"status"`
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
	localNow := carbon.Now(time.Local.String())

	// Check if start_date and end_date are provided
	startDateStr, exists := data.Get("start_date")
	if exists {
		r.StartDate.Carbon = carbon.Parse(startDateStr.(string))
	}

	endDateStr, exists := data.Get("end_date")
	if exists {
		r.EndDate.Carbon = carbon.Parse(endDateStr.(string))
	}

	// Check if start_date and end_date are valid
	if r.StartDate.StdTime().After(r.EndDate.StdTime()) {
		return errors.New("tanggal selesai harus setelah tanggal mulai")
	}

	// Check if start_date is in the future
	if !r.StartDate.StdTime().After(localNow.StdTime()) {
		return errors.New("tanggal mulai harus di masa depan")
	}

	return nil
}
