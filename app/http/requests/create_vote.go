package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateVote struct {
	PollID   string `json:"poll_id"`
	OptionID string `json:"option_id"`
}

func (r *CreateVote) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateVote) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"poll_id":   "required|string",
		"option_id": "required|string",
	}
}

func (r *CreateVote) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateVote) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateVote) PrepareForValidation(ctx http.Context, data validation.Data) error {
	p, _ := data.Get("poll_id")
	o, _ := data.Get("option_id")
	r.PollID = p.(string)
	r.OptionID = o.(string)
	return nil
}
