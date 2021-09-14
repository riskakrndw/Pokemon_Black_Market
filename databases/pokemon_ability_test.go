package databases

import (
	"fmt"
	"project/pbm/config"
	"project/pbm/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mock_ability = models.Ability{
		Name: "blaze",
	}

	mock_pokemon_ability = models.PokemonAbility{
		Name: "blaze",
	}
)

func TestCreatePokemonAbilitySuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonAbility{})
	config.DB.Migrator().AutoMigrate(&models.PokemonAbility{})
	new_abilities := models.Abilities{
		Ability: mock_ability,
	}
	ability, err := CreatePokemonAbility(new_abilities)
	if assert.NoError(t, err) {
		assert.Equal(t, "blaze", ability.Name)
	}
}

func TestCreatePokemonAbilityError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonAbility{})
	new_abilities := models.Abilities{
		Ability: mock_ability,
	}
	_, err := CreatePokemonAbility(new_abilities)
	assert.Error(t, err)
}

func TestCheckPokemonAbilitySuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonAbility{})
	config.DB.Migrator().AutoMigrate(&models.PokemonAbility{})
	new_abilities := models.Abilities{
		Ability: mock_ability,
	}
	if err := config.DB.Save(&mock_pokemon_ability).Error; err != nil {
		t.Error(err)
	}
	new_ability, same, err := CheckPokemonAbility(new_abilities)
	if assert.NoError(t, err) {
		assert.Equal(t, "blaze", new_ability.Name)
		assert.Equal(t, true, same)
	}
}

func TestCheckPokemonAbilityError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonAbility{})
	new_abilities := models.Abilities{
		Ability: mock_ability,
	}
	_, _, err := CheckPokemonAbility(new_abilities)
	assert.Error(t, err)
}

func TestCheckDetailAbilitySuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.PokemonAbility{})
	config.DB.Migrator().AutoMigrate(&models.PokemonAbility{})
	config.DB.Migrator().DropTable(&models.DetailAbility{})
	config.DB.Migrator().AutoMigrate(&models.DetailAbility{})
	if err := config.DB.Save(&mock_pokemon_ability).Error; err != nil {
		t.Error(err)
	}
	mock_pokemon := models.Pokemon{
		Name:   "Riska",
		Weight: 10,
		Height: 10,
		Price:  10000,
		Stock:  50,
	}
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	mock_detail_ability := models.DetailAbility{
		PokemonID:        mock_pokemon.ID,
		PokemonAbilityID: mock_pokemon_ability.ID,
	}
	if err := config.DB.Save(&mock_detail_ability).Error; err != nil {
		t.Error(err)
	}
	new_detail_ability, _, err := CheckDetailAbility(mock_detail_ability.PokemonAbilityID)
	if assert.NoError(t, err) {
		assert.Equal(t, uint(1), new_detail_ability.PokemonAbilityID)
	}
}

func TestCheckDetailAbilityError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.PokemonAbility{})
	config.DB.Migrator().AutoMigrate(&models.PokemonAbility{})
	config.DB.Migrator().DropTable(&models.DetailAbility{})
	if err := config.DB.Save(&mock_pokemon_ability).Error; err != nil {
		t.Error(err)
	}
	mock_pokemon := models.Pokemon{
		Name:   "Riska",
		Weight: 10,
		Height: 10,
		Price:  10000,
		Stock:  50,
	}
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	mock_detail_ability := models.DetailAbility{
		PokemonID:        mock_pokemon.ID,
		PokemonAbilityID: mock_pokemon_ability.ID,
	}
	if err := config.DB.Save(&mock_detail_ability).Error; err != nil {
		t.Error(err)
	}
	_, _, err := CheckDetailAbility(mock_pokemon_ability.ID)
	assert.Error(t, err)
}

func TestCreateDetailAbilitySuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.PokemonAbility{})
	config.DB.Migrator().AutoMigrate(&models.PokemonAbility{})
	config.DB.Migrator().DropTable(&models.DetailAbility{})
	if err := config.DB.Save(&mock_pokemon_ability).Error; err != nil {
		t.Error(err)
	}
	mock_pokemon := models.Pokemon{
		Name:   "Riska",
		Weight: 10,
		Height: 10,
		Price:  10000,
		Stock:  50,
	}
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	detail_ability, err := CreateDetailAbility(mock_pokemon.ID, mock_pokemon_ability.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, uint(1), detail_ability.PokemonAbilityID)
		assert.Equal(t, uint(1), detail_ability.PokemonID)
	}
}

func TestGetPokemonAbilitiesByNameSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonAbility{})
	config.DB.Migrator().AutoMigrate(&models.PokemonAbility{})
	if err := config.DB.Save(&mock_pokemon_ability).Error; err != nil {
		t.Error(err)
	}
	ability, err := GetPokemonAbilitiesByName(mock_pokemon_ability.Name)
	if assert.NoError(t, err) {
		assert.Equal(t, "blaze", ability.Name)
	}
}

func TestGetPokemonAbilitiesByIdSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonAbility{})
	config.DB.Migrator().AutoMigrate(&models.PokemonAbility{})
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.DetailAbility{})
	config.DB.Migrator().AutoMigrate(&models.DetailAbility{})
	if err := config.DB.Save(&mock_pokemon_ability).Error; err != nil {
		t.Error(err)
	}
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	mock_detail_ability := models.DetailAbility{
		PokemonID:        mock_pokemon.ID,
		PokemonAbilityID: mock_pokemon_ability.ID,
	}
	if err := config.DB.Save(&mock_detail_ability).Error; err != nil {
		t.Error(err)
	}
	ability_pokemon, err := GetPokemonAbilitiesById(int(mock_pokemon.ID))
	fmt.Println(ability_pokemon)
	if assert.NoError(t, err) {
		assert.Equal(t, "blaze", ability_pokemon[0].Name)
	}
}
