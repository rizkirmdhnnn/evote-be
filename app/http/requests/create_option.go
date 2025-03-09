package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateOption struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Avatar string `json:"avatar"`
	PollID string `json:"poll_id"`
}

func (r *CreateOption) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateOption) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateOption) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":    "required|string",
		"desc":    "required|string",
		"poll_id": "required|string",
		"avatar":  "string",
	}
}

func (r *CreateOption) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateOption) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateOption) PrepareForValidation(ctx http.Context, data validation.Data) error {
	v, _ := data.Get("poll_id")
	r.PollID = v.(string)
	return nil
}
