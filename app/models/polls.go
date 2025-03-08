package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

// Status Poll enum type
type Status string

const (
	Active Status = "active"
	Done   Status = "done"
)

type Polls struct {
	orm.Model
	Title       string
	Description string
	Status      Status
	StartDate   time.Time
	EndDate     time.Time
	UserID      uint
	orm.SoftDeletes
}

type CreatePollingResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type PollsResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type UpdatePollingResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

func (p *Polls) ToResponse() PollsResponse {
	return PollsResponse{
		ID:          int(p.ID),
		Title:       p.Title,
		Description: p.Description,
		Status:      p.Status,
		StartDate:   p.StartDate.String(),
		EndDate:     p.EndDate.String(),
	}
}
