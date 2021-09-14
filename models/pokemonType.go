package models

type PokemonType struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name" form:"name"`

	//many to many
	Pokemon []*Pokemon `gorm:"many2many:detail_types" json:"detail_types"`
}
