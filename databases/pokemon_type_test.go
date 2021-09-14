package databases

import (
	"fmt"
	"project/pbm/config"
	"project/pbm/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mock_type = models.Type{
		Name: "blaze",
	}

	mock_pokemon_type = models.PokemonType{
		Name: "blaze",
	}
)

func TestCreatePokemonTypeSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonType{})
	config.DB.Migrator().AutoMigrate(&models.PokemonType{})
	new_abilities := models.Types{
		Type: mock_type,
	}
	type_pokemon, err := CreatePokemonType(new_abilities)
	if assert.NoError(t, err) {
		assert.Equal(t, "blaze", type_pokemon.Name)
	}
}

func TestCreatePokemonTypeError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonType{})
	new_abilities := models.Types{
		Type: mock_type,
	}
	_, err := CreatePokemonType(new_abilities)
	assert.Error(t, err)
}

func TestCheckPokemonTypeSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonType{})
	config.DB.Migrator().AutoMigrate(&models.PokemonType{})
	new_abilities := models.Types{
		Type: mock_type,
	}
	if err := config.DB.Save(&mock_pokemon_type).Error; err != nil {
		t.Error(err)
	}
	new_type, same, err := CheckPokemonType(new_abilities)
	if assert.NoError(t, err) {
		assert.Equal(t, "blaze", new_type.Name)
		assert.Equal(t, true, same)
	}
}

func TestCheckPokemonTypeError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonType{})
	new_abilities := models.Types{
		Type: mock_type,
	}
	_, _, err := CheckPokemonType(new_abilities)
	assert.Error(t, err)
}

func TestCheckDetailTypeSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.PokemonType{})
	config.DB.Migrator().AutoMigrate(&models.PokemonType{})
	config.DB.Migrator().DropTable(&models.DetailType{})
	config.DB.Migrator().AutoMigrate(&models.DetailType{})
	if err := config.DB.Save(&mock_pokemon_type).Error; err != nil {
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
	mock_detail_type := models.DetailType{
		PokemonID:     mock_pokemon.ID,
		PokemonTypeID: mock_pokemon_type.ID,
	}
	if err := config.DB.Save(&mock_detail_type).Error; err != nil {
		t.Error(err)
	}
	new_detail_type, _, err := CheckDetailType(mock_detail_type.PokemonTypeID)
	if assert.NoError(t, err) {
		assert.Equal(t, uint(1), new_detail_type.PokemonTypeID)
	}
}

func TestCheckDetailTypeError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.PokemonType{})
	config.DB.Migrator().AutoMigrate(&models.PokemonType{})
	config.DB.Migrator().DropTable(&models.DetailType{})
	if err := config.DB.Save(&mock_pokemon_type).Error; err != nil {
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
	mock_detail_type := models.DetailType{
		PokemonID:     mock_pokemon.ID,
		PokemonTypeID: mock_pokemon_type.ID,
	}
	if err := config.DB.Save(&mock_detail_type).Error; err != nil {
		t.Error(err)
	}
	_, _, err := CheckDetailType(mock_pokemon_type.ID)
	assert.Error(t, err)
}

func TestCreateDetailTypeSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.PokemonType{})
	config.DB.Migrator().AutoMigrate(&models.PokemonType{})
	config.DB.Migrator().DropTable(&models.DetailType{})
	if err := config.DB.Save(&mock_pokemon_type).Error; err != nil {
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
	detail_type, err := CreateDetailType(mock_pokemon.ID, mock_pokemon_type.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, uint(1), detail_type.PokemonTypeID)
		assert.Equal(t, uint(1), detail_type.PokemonID)
	}
}

func TestGetPokemonTypesByNameSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonType{})
	config.DB.Migrator().AutoMigrate(&models.PokemonType{})
	if err := config.DB.Save(&mock_pokemon_type).Error; err != nil {
		t.Error(err)
	}
	type_pokemon, err := GetPokemonTypesByName(mock_pokemon_type.Name)
	if assert.NoError(t, err) {
		assert.Equal(t, "blaze", type_pokemon.Name)
	}
}

func TestGetPokemonTypesByIdSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.PokemonType{})
	config.DB.Migrator().AutoMigrate(&models.PokemonType{})
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.DetailType{})
	config.DB.Migrator().AutoMigrate(&models.DetailType{})
	if err := config.DB.Save(&mock_pokemon_type).Error; err != nil {
		t.Error(err)
	}
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	mock_detail_type := models.DetailType{
		PokemonID:     mock_pokemon.ID,
		PokemonTypeID: mock_pokemon_type.ID,
	}
	if err := config.DB.Save(&mock_detail_type).Error; err != nil {
		t.Error(err)
	}
	type_pokemon, err := GetPokemonTypesById(int(mock_pokemon.ID))
	fmt.Println(type_pokemon)
	if assert.NoError(t, err) {
		assert.Equal(t, "blaze", type_pokemon[0].Name)
	}
}
