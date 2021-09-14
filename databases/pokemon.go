package databases

import (
	"project/pbm/config"
	"project/pbm/models"
)

func CheckPokemon(name string) (models.Pokemon, bool, error) {
	var pokemon models.Pokemon

	if err := config.DB.Where("name = ?", name).First(&pokemon).Error; err != nil {
		return pokemon, false, err
	}

	if pokemon.ID == 0 {
		return pokemon, false, nil
	} else {
		return pokemon, true, nil
	}
}

func CheckPokemonDeleted(name string) (models.Pokemon, bool, error) {
	var pokemon models.Pokemon

	if err := config.DB.Raw("SELECT * FROM pokemons WHERE name = ? AND deleted_at IS NOT NULL", name).Scan(&pokemon).Error; err != nil {
		return pokemon, false, err
	}

	if pokemon.ID == 0 {
		return pokemon, false, nil
	} else {
		return pokemon, true, nil
	}
}

func CreatePokemon(pokemon models.Pokemon) (models.Pokemon, error) {
	if err := config.DB.Save(&pokemon).Error; err != nil {
		return pokemon, err
	}

	return pokemon, nil
}

func GetAllPokemon() ([]models.Pokemon, error) {
	var all_pokemon []models.Pokemon
	if err := config.DB.Find(&all_pokemon).Error; err != nil {
		return all_pokemon, err
	}
	return all_pokemon, nil
}

func GetPokemonById(id int) (models.Pokemon, error) {
	var pokemon models.Pokemon

	if err := config.DB.Where("id = ?", id).First(&pokemon).Error; err != nil {
		return pokemon, err
	}

	return pokemon, nil
}

func GetPokemonByName(name string) ([]models.Pokemon, error) {
	var pokemons []models.Pokemon
	search := "%" + name + "%"
	if err := config.DB.Find(&pokemons, "name LIKE ?", search).Error; err != nil {
		return pokemons, err
	}

	return pokemons, nil
}

func UpdatePokemon(pokemon models.Pokemon) (models.Pokemon, error) {
	if err := config.DB.Save(&pokemon).Error; err != nil {
		return pokemon, err
	}

	return pokemon, nil
}

func UpdatePokemonDeleted(pokemon models.Pokemon, id uint) (models.Pokemon, error) {
	if err := config.DB.Raw("UPDATE pokemons SET deleted_at = NULL WHERE id = ?", id).Scan(&pokemon).Error; err != nil {
		return pokemon, err
	}

	return pokemon, nil
}

func DeletePokemon(id int) (models.Pokemon, error) {
	var pokemon models.Pokemon
	if err := config.DB.Find(&pokemon, "id = ?", id).Error; err != nil {
		return pokemon, err
	}
	if err := config.DB.Delete(&pokemon, "id = ?", id).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}
