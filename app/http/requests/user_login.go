package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *UserLogin) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserLogin) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserLogin) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email":    "required|email",
		"password": "required|min_len:6",
	}
}

func (r *UserLogin) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserLogin) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserLogin) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
