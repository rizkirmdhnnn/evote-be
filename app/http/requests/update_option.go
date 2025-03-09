package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdateOption struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Avatar string `json:"avatar"`
	PollID string `json:"poll_id"`
}

func (r *UpdateOption) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateOption) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateOption) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":    "string",
		"desc":    "string",
		"avatar":  "string",
		"poll_id": "required|string",
	}
}

func (r *UpdateOption) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateOption) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateOption) PrepareForValidation(ctx http.Context, data validation.Data) error {
	v, _ := data.Get("poll_id")
	r.PollID = v.(string)
	return nil
}
