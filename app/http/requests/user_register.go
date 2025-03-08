package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *UserRegister) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserRegister) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserRegister) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":     "required|string|min_len:3",
		"email":    "required|email",
		"password": "required|string|min_len:6",
	}
}

func (r *UserRegister) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserRegister) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserRegister) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
