package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name              string
	Email             string
	Password          string
	Avatar            string
	EmailVerifiedAt  string
	VerificationToken string
	Polls             []*Polls
	orm.SoftDeletes
}

type UserRegisterResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type UserLoginResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	Token  string `json:"token"`
}
