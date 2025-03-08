package requests

import (
	"errors"
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreatePolling struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	StartDate   carbon.DateTime `json:"start_date"`
	EndDate     carbon.DateTime `json:"end_date"`
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
	localNow := carbon.Now(time.Local.String())

	// Check if start_date and end_date are provided
	startDateStr, exists := data.Get("start_date")
	if !exists {
		return errors.New("start_date is required")
	}
	endDateStr, exists := data.Get("end_date")
	if !exists {
		return errors.New("end_date is required")
	}

	// Parse start_date dan end_date tanpa mengubah timezone
	startDate := carbon.Parse(startDateStr.(string))
	endDate := carbon.Parse(endDateStr.(string))

	// Simpan hasil parsing ke struct
	r.StartDate.Carbon = startDate
	r.EndDate.Carbon = endDate

	// Cek jika start_date lebih besar dari end_date
	if startDate.StdTime().After(endDate.StdTime()) {
		return errors.New("tanggal selesai harus setelah tanggal mulai")
	}

	// Cek jika start_date berada di masa depan
	if !startDate.StdTime().After(localNow.StdTime()) {
		return errors.New("tanggal mulai harus di masa depan")
	}

	return nil
}
