package models

type PokemonAbility struct {
	ID   uint   `gorm:"primary_key" json:"id" form:"id"`
	Name string `json:"name" form:"name"`

	//many to many
	Pokemon []*Pokemon `gorm:"many2many:detail_abilities" json:"detail_abilities"`
}
