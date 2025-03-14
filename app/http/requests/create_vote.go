package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateVote struct {
	Code     string `json:"code"`
	OptionID string `json:"option_id"`
}

func (r *CreateVote) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateVote) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"code":      "required|string",
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
	value, isExists := data.Get("option_id")
	if !isExists {
		return nil
	}
	r.OptionID = value.(string)
	return nil
}
