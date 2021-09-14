package databases

import (
	"project/pbm/config"
	"project/pbm/models"
)

func CheckPokemonType(types models.Types) (models.PokemonType, bool, error) {
	var pokemon_type models.PokemonType

	if err := config.DB.Model(&pokemon_type).Where("name = ?", types.Type.Name).First(&pokemon_type).Error; err != nil {
		return pokemon_type, false, err
	}
	if pokemon_type.Name == types.Type.Name {
		return pokemon_type, true, nil
	} else {
		return pokemon_type, false, nil
	}
}

func CreatePokemonType(type_pokemon models.Types) (models.PokemonType, error) {
	pokemon_type := models.PokemonType{
		Name: type_pokemon.Type.Name,
	}
	if err := config.DB.Save(&pokemon_type).Error; err != nil {
		return pokemon_type, err
	}
	return pokemon_type, nil
}

func CheckDetailType(type_id uint) (models.DetailType, bool, error) {
	var detail_type models.DetailType

	if err := config.DB.Where("pokemon_type_id = ?", type_id).First(&detail_type).Error; err != nil {
		return detail_type, true, err
	}
	return detail_type, false, nil
}

func CreateDetailType(pokemon_id, pokemon_type_id uint) (models.DetailType, error) {
	detail_type := models.DetailType{
		PokemonID:     pokemon_id,
		PokemonTypeID: pokemon_type_id,
	}
	if err := config.DB.Save(&detail_type).Error; err != nil {
		return detail_type, err
	}
	return detail_type, nil
}

func GetPokemonTypesByName(pokemon_type_name string) (models.PokemonType, error) {
	var pokemon_type models.PokemonType

	if err := config.DB.Find(&pokemon_type, "name = ?", pokemon_type_name).Error; err != nil {
		return pokemon_type, err
	}
	return pokemon_type, nil
}

func GetPokemonTypesById(pokemon_id int) ([]models.PokemonType, error) {
	var pokemon_type []models.PokemonType

	if err := config.DB.Model(&pokemon_type).Joins("JOIN detail_types on detail_types.pokemon_type_id = pokemon_types.id").Joins("JOIN pokemons on pokemons.id = detail_types.pokemon_id").Where("pokemons.id = ?", pokemon_id).Scan(&pokemon_type).Error; err != nil {
		return pokemon_type, err
	}
	return pokemon_type, nil
}

// func GetPokemonTypesByPokemonId(pokemon_id int) (models.PokemonType, error) {
// 	var pokemon_type models.PokemonType
// 	if err := config.DB.Model(&pokemon_type).Joins("JOIN detail_types on detail_types.pokemon_type_id = pokemon_types.id").Joins("JOIN pokemons on pokemons.id = detail_types.pokemon_id").Where("pokemons.id = ?", pokemon_id).Scan(&pokemon_type).Error; err != nil {
// 		return pokemon_type, err
// 	}
// 	return pokemon_type, nil
// }

// func CreatePokemonType(types []models.Types) []models.PokemonType {
// 	var pokemon_type models.PokemonType
// 	var pokemon_types []models.PokemonType
// 	for _, v := range types {
// 		pokemon_type = models.PokemonType{
// 			Name: v.Type.Name,
// 		}
// 		config.DB.Raw("INSERT INTO pokemon_types (name) SELECT * FROM (SELECT ?) AS tmp WHERE NOT EXISTS (SELECT name FROM pokemon_types WHERE name = ?) LIMIT 1", pokemon_type.Name, pokemon_type.Name).Scan(&pokemon_types)

// 		for i := 0; i < len(types); i++ {
// 			new_array := models.PokemonType{
// 				Name: types[i].Type.Name,
// 			}
// 			pokemon_types = append(pokemon_types, new_array)
// 		}
// 	}
// 	return pokemon_types
// }
