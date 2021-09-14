package models

type DetailType struct {
	PokemonID     uint `gorm:"primaryKey" json:"pokemon_id" form:"pokemon_id"`
	PokemonTypeID uint `gorm:"primaryKey" json:"type_id" form:"type_id"`
}
