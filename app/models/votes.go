package models

import (
	"github.com/goravel/framework/database/orm"
)

type Votes struct {
	orm.Model
	UserID   uint
	PollID   uint
	OptionID uint
	Polls    Polls `gorm:"foreignKey:PollID"`
	orm.SoftDeletes
}
