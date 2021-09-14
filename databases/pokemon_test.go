package databases

import (
	"project/pbm/config"
	"project/pbm/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mock_pokemon = models.Pokemon{
		Name:   "bulbasaur",
		Weight: 10,
		Height: 20,
		Price:  500000,
		Stock:  100,
	}
)

func TestCreatePokemonSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	pokemon, err := CreatePokemon(mock_pokemon)
	if assert.NoError(t, err) {
		assert.Equal(t, "bulbasaur", pokemon.Name)
		assert.Equal(t, 10, pokemon.Weight)
		assert.Equal(t, 20, pokemon.Height)
		assert.Equal(t, 500000, pokemon.Price)
		assert.Equal(t, 100, pokemon.Stock)
	}
}

func TestCreatePokemonError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	_, err := CreatePokemon(mock_pokemon)
	assert.Error(t, err)
}

func TestGetAllPokemonSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	pokemon, err := GetAllPokemon()
	if assert.NoError(t, err) {
		assert.Equal(t, "bulbasaur", pokemon[0].Name)
		assert.Equal(t, 10, pokemon[0].Weight)
		assert.Equal(t, 20, pokemon[0].Height)
		assert.Equal(t, 500000, pokemon[0].Price)
		assert.Equal(t, 100, pokemon[0].Stock)
	}
}

func TestGetAllPokemonError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	CreatePokemon(mock_pokemon)
	_, err := GetAllPokemon()
	assert.Error(t, err)
}

func TestGetPokemonByIdSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	pokemon, err := GetPokemonById(int(mock_pokemon.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "bulbasaur", pokemon.Name)
		assert.Equal(t, 10, pokemon.Weight)
		assert.Equal(t, 20, pokemon.Height)
		assert.Equal(t, 500000, pokemon.Price)
		assert.Equal(t, 100, pokemon.Stock)
	}
}

func TestGetPokemonByIdError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	CreatePokemon(mock_pokemon)
	_, err := GetPokemonById(int(mock_pokemon.ID))
	assert.Error(t, err)
}

func TestGetPokemonByNameSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	pokemon, err := GetPokemonByName(mock_pokemon.Name)
	if assert.NoError(t, err) {
		assert.Equal(t, "bulbasaur", pokemon[0].Name)
		assert.Equal(t, 10, pokemon[0].Weight)
		assert.Equal(t, 20, pokemon[0].Height)
		assert.Equal(t, 500000, pokemon[0].Price)
		assert.Equal(t, 100, pokemon[0].Stock)
	}
}

func TestGetPokemonByNameError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	CreatePokemon(mock_pokemon)
	_, err := GetPokemonByName(mock_pokemon.Name)
	assert.Error(t, err)
}

func TestUpdatePokemonSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	mock_pokemon.Stock = 500
	pokemon, err := UpdatePokemon(mock_pokemon)
	if assert.NoError(t, err) {
		assert.Equal(t, "bulbasaur", pokemon.Name)
		assert.Equal(t, 10, pokemon.Weight)
		assert.Equal(t, 20, pokemon.Height)
		assert.Equal(t, 500000, pokemon.Price)
		assert.Equal(t, 500, pokemon.Stock)
	}
}

func TestUpdatePokemonError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	CreatePokemon(mock_pokemon)
	mock_pokemon.Stock = 500
	_, err := UpdatePokemon(mock_pokemon)
	assert.Error(t, err)
}

func TestUpdatePokemonDeletedSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	pokemon, err := UpdatePokemonDeleted(mock_pokemon, mock_pokemon.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, "bulbasaur", pokemon.Name)
		assert.Equal(t, 10, pokemon.Weight)
		assert.Equal(t, 20, pokemon.Height)
		assert.Equal(t, 500000, pokemon.Price)
		assert.Equal(t, 100, pokemon.Stock)
	}
}

func TestUpdatePokemonDeletedError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	CreatePokemon(mock_pokemon)
	_, err := UpdatePokemonDeleted(mock_pokemon, mock_pokemon.ID)
	assert.Error(t, err)
}

func TestDeletePokemonSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	pokemon, err := DeletePokemon(int(mock_pokemon.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "bulbasaur", pokemon.Name)
		assert.Equal(t, 10, pokemon.Weight)
		assert.Equal(t, 20, pokemon.Height)
		assert.Equal(t, 500000, pokemon.Price)
		assert.Equal(t, 100, pokemon.Stock)
	}
}

func TestDeletePokemonError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	CreatePokemon(mock_pokemon)
	_, err := DeletePokemon(int(mock_pokemon.ID))
	assert.Error(t, err)
}

func TestCheckPokemonSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	pokemon, same, err := CheckPokemon(mock_pokemon.Name)
	if assert.NoError(t, err) {
		assert.Equal(t, true, same)
		assert.Equal(t, "bulbasaur", pokemon.Name)
	}
}

func TestCheckPokemonError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	CreatePokemon(mock_pokemon)
	_, _, err := CheckPokemon(mock_pokemon.Name)
	assert.Error(t, err)
}

func TestCheckPokemonDeletedSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	if err := config.DB.Save(&mock_pokemon).Error; err != nil {
		t.Error(err)
	}
	pokemon, same, err := CheckPokemonDeleted(mock_pokemon.Name)
	if assert.NoError(t, err) {
		assert.Equal(t, false, same)
		assert.Equal(t, "", pokemon.Name)
	}
}

func TestCheckPokemonDeletedError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.Pokemon{})
	CreatePokemon(mock_pokemon)
	_, _, err := CheckPokemonDeleted(mock_pokemon.Name)
	assert.Error(t, err)
}
