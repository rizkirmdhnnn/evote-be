package requests

import (
	"mime/multipart"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateOption struct {
	Name   string               `json:"name" form:"name"`
	Desc   string               `json:"desc" form:"desc"`
	Avatar multipart.FileHeader `json:"avatar" form:"avatar" swaggerignore:"true"`
	PollID string               `json:"poll_id" form:"poll_id"`
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
		"avatar":  "file|image",
	}
}

func (r *CreateOption) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateOption) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateOption) PrepareForValidation(ctx http.Context, data validation.Data) error {
	value, isExists := data.Get("poll_id")
	if !isExists {
		return nil
	}
	r.PollID = value.(string)
	return nil
}
