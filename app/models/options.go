package models

import (
	"github.com/goravel/framework/database/orm"
)

type Options struct {
	orm.Model
	Name       string
	Desc       string
	Avatar     string
	PollID     uint
	VotesCount uint
	Votes      []*Votes `gorm:"foreignKey:OptionID"`
	orm.SoftDeletes
}

type CreateOptionsResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Avatar     string `json:"avatar"`
	VotesCount uint   `json:"votes_count"`
}

func (r *Options) ToResponse() CreateOptionsResponse {
	return CreateOptionsResponse{
		ID:         int(r.ID),
		Name:       r.Name,
		Desc:       r.Desc,
		Avatar:     r.Avatar,
		VotesCount: r.VotesCount,
	}
}
