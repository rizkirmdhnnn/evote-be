package requests

import (
	"mime/multipart"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdateOption struct {
	Name   string               `json:"name" form:"name"`
	Desc   string               `json:"desc" form:"desc"`
	Avatar multipart.FileHeader `json:"avatar" form:"avatar" swaggerignore:"true"`
	PollID string               `json:"poll_id" form:"poll_id"`
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
		"avatar":  "file|image",
		"poll_id": "string",
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
