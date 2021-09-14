package models

type DetailAbility struct {
	PokemonID        uint `gorm:"primaryKey" json:"pokemon_id" form:"pokemon_id"`
	PokemonAbilityID uint `gorm:"primaryKey" json:"ability_id" form:"ability_id"`
}
