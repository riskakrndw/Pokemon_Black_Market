package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Token    string `json:"token" form:"token"`

	//FK
	LevelID uint `json:"level_id" form:"level_id"`

	//1 to many
	Transaction []Transaction `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
}
