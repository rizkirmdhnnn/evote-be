package requests

import (
	"errors"
	"fmt"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreatePolling struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StartDate   *time.Time `json:"start_date" swaggertype:"string" example:"2022-01-01 00:00" format:"date-time"`
	EndDate     time.Time  `json:"end_date" swaggertype:"string" example:"2022-01-01 00:00" format:"date-time"`
	// testing:
	// * active - Active, can be voted
	// * done - Done, can't be voted
	Status string `json:"status" swaggertype:"string" enums:"Active,Done,Scheduled"`
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
		"end_date":    "required|date",
	}
}

func (r *CreatePolling) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreatePolling) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreatePolling) PrepareForValidation(ctx http.Context, data validation.Data) error {
	// Dapatkan waktu lokal saat ini
	localNow := time.Now().Truncate(time.Minute)

	// Parse tanggal dari string
	layout := "2006-01-02 15:04"
	startDateStr, exists := data.Get("start_date")
	if !exists {
		startDateStr = localNow.Format(layout)
	}
	endDateStr, exists := data.Get("end_date")
	if !exists {
		return errors.New("end_date is required")
	}

	// Parse dengan zona waktu yang sama
	startDate, err := time.ParseInLocation(layout, startDateStr.(string), localNow.Location())
	if err != nil {
		return errors.New("invalid start_date format, use YYYY-MM-DD HH:MM format")
	}

	endDate, err := time.ParseInLocation(layout, endDateStr.(string), localNow.Location())
	if err != nil {
		return errors.New("invalid end_date format, use YYYY-MM-DD HH:MM format")
	}

	// Simpan ke struct
	r.StartDate = &startDate
	r.EndDate = endDate

	// Debug
	fmt.Println("Local Now:", localNow)
	fmt.Println("Start Date:", startDate)
	fmt.Println("End Date:", endDate)

	// Validasi tanggal mulai harus di masa depan
	if startDate.Before(localNow) {
		return errors.New("invalid start_date, must be in the future")
	}

	// Validasi tanggal mulai harus sebelum tanggal selesai
	if startDate.After(endDate) || startDate.Equal(endDate) {
		return errors.New("invalid date range, start_date must be before end_date")
	}

	return nil
}
