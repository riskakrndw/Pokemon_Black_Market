package models

import "gorm.io/gorm"

type Pokemon struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	Weight int    `json:"weight" form:"weight"`
	Height int    `json:"height" form:"height"`
	Price  int    `json:"price" form:"price"`
	Stock  int    `json:"stock" form:"stock"`

	//many to many
	PokemonType    []*PokemonType    `gorm:"many2many:detail_types" json:"detail_types"`
	PokemonAbility []*PokemonAbility `gorm:"many2many:detail_abilities" json:"detail_abilities"`
	Transaction    []*Transaction    `gorm:"many2many:detail_transactions" json:"detail_transactions"`
}
