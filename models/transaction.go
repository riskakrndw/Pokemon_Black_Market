package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Code        string    `json:"code" form:"code"`
	DateOfMonth time.Time `json:"dom" form:"dom"`
	Total       int       `json:"total" form:"total"`
	Status      string    `json:"status" form:"status"`

	//FK
	UserID uint `json:"user_id" form:"user_id"`

	//many to many
	Pokemon []*Pokemon `gorm:"many2many:detail_transactions" json:"detail_transactions"`
}
