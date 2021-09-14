package databases

import (
	"project/pbm/config"
	"project/pbm/models"
)

func CheckPokemonAbility(ability models.Abilities) (models.PokemonAbility, bool, error) {
	var pokemon_ability models.PokemonAbility

	if err := config.DB.Model(&pokemon_ability).Where("name=?", ability.Ability.Name).First(&pokemon_ability).Error; err != nil {
		return pokemon_ability, false, err
	}
	if pokemon_ability.Name == ability.Ability.Name {
		return pokemon_ability, true, nil
	} else {
		return pokemon_ability, false, nil
	}
}

func CreatePokemonAbility(ability_pokemon models.Abilities) (models.PokemonAbility, error) {
	pokemon_ability := models.PokemonAbility{
		Name: ability_pokemon.Ability.Name,
	}
	if err := config.DB.Save(&pokemon_ability).Error; err != nil {
		return pokemon_ability, err
	}
	return pokemon_ability, nil
}

func CheckDetailAbility(ability_id uint) (models.DetailAbility, bool, error) {
	var detail_ability models.DetailAbility

	if err := config.DB.Where("pokemon_ability_id = ?", ability_id).First(&detail_ability).Error; err != nil {
		return detail_ability, true, err
	}
	return detail_ability, false, nil
}

func CreateDetailAbility(pokemon_id, pokemon_ability_id uint) (models.DetailAbility, error) {
	detail_ability := models.DetailAbility{
		PokemonID:        pokemon_id,
		PokemonAbilityID: pokemon_ability_id,
	}
	if err := config.DB.Save(&detail_ability).Error; err != nil {
		return detail_ability, err
	}
	return detail_ability, nil
}

func GetPokemonAbilitiesByName(pokemon_ability_name string) (models.PokemonAbility, error) {
	var pokemon_ability models.PokemonAbility

	if err := config.DB.Find(&pokemon_ability, "name = ?", pokemon_ability_name).Error; err != nil {
		return pokemon_ability, err
	}
	return pokemon_ability, nil
}

func GetPokemonAbilitiesById(pokemon_id int) ([]models.PokemonAbility, error) {
	var pokemon_abilities []models.PokemonAbility

	if err := config.DB.Model(&pokemon_abilities).Joins("JOIN detail_abilities on detail_abilities.pokemon_ability_id = pokemon_abilities.id").Joins("JOIN pokemons on pokemons.id = detail_abilities.pokemon_id").Where("pokemons.id = ?", pokemon_id).Scan(&pokemon_abilities).Error; err != nil {
		return pokemon_abilities, err
	}
	return pokemon_abilities, nil
}
