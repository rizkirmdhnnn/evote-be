package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserUpdate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r *UserUpdate) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserUpdate) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserUpdate) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":  "string",
		"email": "email",
	}
}

func (r *UserUpdate) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserUpdate) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserUpdate) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
