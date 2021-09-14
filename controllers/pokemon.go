package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"project/pbm/databases"
	"project/pbm/middlewares"
	"project/pbm/models"
	"strconv"

	"github.com/labstack/echo"
)

type Types struct {
	Name string
}

type Abilities struct {
	Name string
}

type Output1 struct {
	Name   string
	ID     uint
	Weight int
	Height int
	Price  int
	Stock  int
}

type Output2 struct {
	Name      string
	ID        uint
	Weight    int
	Height    int
	Price     int
	Stock     int
	Types     []models.PokemonType
	Abilities []models.PokemonAbility
}

type OutputSearch struct {
	Name      string
	ID        int
	Weight    int
	Height    int
	Types     []Types
	Abilities []Abilities
}

func SearchAllPokemon(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID != 2 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Operasional")
	}

	//get pokemon from api
	response, _ := http.Get("https://pokeapi.co/api/v2/pokemon")
	response_data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var response_object models.Response
	json.Unmarshal(response_data, &response_object)

	//set pokemon into struct
	var pokemon []models.SearchPokemon
	for i := 0; i < len(response_object.Results); i++ {
		detail := GetDetailAllPokemon(response_object.Results[i].Url)
		new_result := models.SearchPokemon{
			ID:        detail.ID,
			Url:       response_object.Results[i].Url,
			Name:      detail.Name,
			Weight:    detail.Weight,
			Height:    detail.Height,
			Types:     detail.Types,
			Abilities: detail.Abilities,
		}
		pokemon = append(pokemon, new_result)
	}

	return c.JSON(http.StatusOK, pokemon)
}

func GetDetailAllPokemon(url string) models.SearchPokemon {
	var pokemon models.SearchPokemon
	response, _ := http.Get(url)
	response_data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	json.Unmarshal(response_data, &pokemon)

	return pokemon
}

func SearchPokemon(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID != 2 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Operasional")
	}

	//get param
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "name cannot null",
		})
	}

	//get pokemon from api
	response, _ := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)
	response_data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var response_object models.SearchPokemon
	if err := json.Unmarshal(response_data, &response_object); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "cannot find pokemon",
		})
	}

	//customize output
	var types []Types
	for i := 0; i < len(response_object.Types); i++ {
		new_result := Types{
			Name: response_object.Types[i].Type.Name,
		}
		types = append(types, new_result)
	}
	var abilities []Abilities
	for i := 0; i < len(response_object.Abilities); i++ {
		new_result := Abilities{
			Name: response_object.Abilities[i].Ability.Name,
		}
		abilities = append(abilities, new_result)
	}

	//set output data
	output := OutputSearch{
		Name:      response_object.Name,
		ID:        response_object.ID,
		Weight:    response_object.Weight,
		Height:    response_object.Height,
		Types:     types,
		Abilities: abilities,
	}

	return c.JSON(http.StatusOK, output)
}

func CreatePokemon(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID != 2 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Operasional")
	}

	//get query param
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "name cannot null",
		})
	}

	//search pokemon from api
	response, _ := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)
	response_data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var response_object models.SearchPokemon
	if err := json.Unmarshal(response_data, &response_object); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "cannot find pokemon",
		})
	}

	//create pokemon
	input_user := models.Pokemon{}
	c.Bind(&input_user)
	input_user = models.Pokemon{
		Name:   response_object.Name,
		Weight: response_object.Weight,
		Height: response_object.Height,
		Price:  input_user.Price,
		Stock:  input_user.Stock,
	}
	pokemon, is_pokemon_exist, _ := databases.CheckPokemon(input_user.Name)
	pokemon_deleted, is_pokemon_deleted, _ := databases.CheckPokemonDeleted(input_user.Name)
	if !is_pokemon_exist && is_pokemon_deleted { //if pokemon deleted, set deleted_at into NULL
		pokemon, err = databases.UpdatePokemonDeleted(pokemon_deleted, pokemon_deleted.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot update data")
		}
	} else if !is_pokemon_exist && !is_pokemon_deleted { //if pokemon not exist and pokemon not deleted
		pokemon, err = databases.CreatePokemon(input_user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot insert data")
		}
	} else if is_pokemon_exist && !is_pokemon_deleted {
		return c.JSON(http.StatusBadRequest, "pokemon already exists")
	}

	//create type pokemon
	var pokemon_types []models.PokemonType
	var pokemon_type models.PokemonType
	for i := 0; i < len(response_object.Types); i++ {
		check_type, is_type_pokemon_exist, _ := databases.CheckPokemonType(response_object.Types[i])
		if is_type_pokemon_exist == false {
			pokemon_type, err = databases.CreatePokemonType(response_object.Types[i])
			if err != nil {
				return c.JSON(http.StatusInternalServerError, "cannot insert data")
			}
		} else {
			pokemon_type = check_type
		}
		pokemon_types = append(pokemon_types, pokemon_type)
	}

	//create ability pokemon
	var pokemon_abilities []models.PokemonAbility
	var pokemon_ability models.PokemonAbility
	for i := 0; i < len(response_object.Abilities); i++ {
		check_ability, is_ability_pokemon_exist, _ := databases.CheckPokemonAbility(response_object.Abilities[i])
		if is_ability_pokemon_exist == false {
			pokemon_ability, err = databases.CreatePokemonAbility(response_object.Abilities[i])
			if err != nil {
				return c.JSON(http.StatusInternalServerError, "cannot insert data")
			}
		} else {
			pokemon_ability = check_ability
		}
		pokemon_abilities = append(pokemon_abilities, pokemon_ability)
	}

	//get pokemon types
	var data_pokemon_types []models.PokemonType
	var get_pokemon_type models.PokemonType
	for i := 0; i < len(pokemon_types); i++ {
		get_pokemon_type, err = databases.GetPokemonTypesByName(pokemon_types[i].Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot get data")
		}
		data_pokemon_types = append(data_pokemon_types, get_pokemon_type)
	}

	//create detail types
	var detail_types []models.DetailType
	var detail_type models.DetailType
	for i := 0; i < len(data_pokemon_types); i++ {
		detail_type, err = databases.CreateDetailType(pokemon.ID, data_pokemon_types[i].ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot insert data")
		}
		detail_types = append(detail_types, detail_type)
	}

	//get pokemon ability
	var data_pokemon_abilities []models.PokemonAbility
	var get_pokemon_ability models.PokemonAbility
	for i := 0; i < len(pokemon_abilities); i++ {
		get_pokemon_ability, err = databases.GetPokemonAbilitiesByName(pokemon_abilities[i].Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot get data")
		}
		data_pokemon_abilities = append(data_pokemon_abilities, get_pokemon_ability)
	}

	//create detail abilities
	var detail_abilities []models.DetailAbility
	var detail_ability models.DetailAbility
	for i := 0; i < len(data_pokemon_abilities); i++ {
		detail_ability, err = databases.CreateDetailAbility(pokemon.ID, data_pokemon_abilities[i].ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot insert data")
		}
		detail_abilities = append(detail_abilities, detail_ability)
	}

	//customize output
	output := Output2{
		Name:      pokemon.Name,
		ID:        pokemon.ID,
		Weight:    pokemon.Weight,
		Height:    pokemon.Height,
		Price:     pokemon.Price,
		Stock:     pokemon.Stock,
		Types:     data_pokemon_types,
		Abilities: data_pokemon_abilities,
	}

	return c.JSON(http.StatusOK, output)
}

func GetAllPokemon(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID == 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Operasional and Pengedar")
	}

	all_pokemon, err := databases.GetAllPokemon()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//customize output
	var output []Output1
	for i := 0; i < len(all_pokemon); i++ {
		new_result := Output1{
			Name:   all_pokemon[i].Name,
			ID:     all_pokemon[i].ID,
			Weight: all_pokemon[i].Weight,
			Height: all_pokemon[i].Height,
			Price:  all_pokemon[i].Price,
			Stock:  all_pokemon[i].Stock,
		}
		output = append(output, new_result)
	}

	return c.JSON(http.StatusOK, output)
}

func GetAllPokemonTesting() echo.HandlerFunc {
	return GetAllPokemon
}

func GetPokemonById(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID != 2 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Operasional")
	}

	//get param
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//get pokemon
	pokemon, err := databases.GetPokemonById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//get detail types
	detail_types, err := databases.GetPokemonTypesById(int(pokemon.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//get pokemon types
	var data_pokemon_types []models.PokemonType
	var get_pokemon_type models.PokemonType
	for i := 0; i < len(detail_types); i++ {
		get_pokemon_type, err = databases.GetPokemonTypesByName(detail_types[i].Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot get data")
		}
		data_pokemon_types = append(data_pokemon_types, get_pokemon_type)
	}

	//get detail abilities
	detail_abilities, err := databases.GetPokemonAbilitiesById(int(pokemon.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//get pokemon abilities
	var data_pokemon_abilities []models.PokemonAbility
	var get_pokemon_ability models.PokemonAbility
	for i := 0; i < len(detail_abilities); i++ {
		get_pokemon_ability, err = databases.GetPokemonAbilitiesByName(detail_abilities[i].Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot get data")
		}
		data_pokemon_abilities = append(data_pokemon_abilities, get_pokemon_ability)
	}

	//customize output
	output := Output2{
		Name:      pokemon.Name,
		ID:        pokemon.ID,
		Weight:    pokemon.Weight,
		Height:    pokemon.Height,
		Price:     pokemon.Price,
		Stock:     pokemon.Stock,
		Types:     data_pokemon_types,
		Abilities: data_pokemon_abilities,
	}

	return c.JSON(http.StatusOK, output)
}

func GetPokemonByIdTesting() echo.HandlerFunc {
	return GetPokemonById
}

func GetPokemonByName(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID == 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Operasional and Pengedar")
	}

	//get param
	name := c.QueryParam("name")

	//get pokemon
	pokemons, err := databases.GetPokemonByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//customize output
	var output []Output1
	for i := 0; i < len(pokemons); i++ {
		new_result := Output1{
			Name:   pokemons[i].Name,
			ID:     pokemons[i].ID,
			Weight: pokemons[i].Weight,
			Height: pokemons[i].Height,
			Price:  pokemons[i].Price,
			Stock:  pokemons[i].Stock,
		}
		output = append(output, new_result)
	}

	return c.JSON(http.StatusOK, output)
}

func GetPokemonByNameTesting() echo.HandlerFunc {
	return GetPokemonByName
}

func UpdatePokemon(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID != 2 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Operasional")
	}

	//get param
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//get old pokemon
	old_pokemon, err := databases.GetPokemonById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//get user's input
	new_pokemon, err := databases.GetPokemonById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	c.Bind(&new_pokemon)

	//check
	if old_pokemon.Name != new_pokemon.Name || old_pokemon.ID != new_pokemon.ID || old_pokemon.Weight != new_pokemon.Weight || old_pokemon.Height != new_pokemon.Height {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Unauthorized update this data",
		})
	}

	//update data
	pokemon, err := databases.UpdatePokemon(new_pokemon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot update data")
	}

	//customize output
	output := Output1{
		Name:   pokemon.Name,
		ID:     pokemon.ID,
		Weight: pokemon.Weight,
		Height: pokemon.Height,
		Price:  pokemon.Price,
		Stock:  pokemon.Stock,
	}

	return c.JSON(http.StatusOK, output)
}

func DeletePokemon(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID != 2 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Operasional")
	}

	//get param
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//delete
	pokemon, err := databases.DeletePokemon(id)
	if pokemon.ID == 0 {
		return c.JSON(http.StatusInternalServerError, "cannot delete data")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//customize output
	output := Output1{
		Name:   pokemon.Name,
		ID:     pokemon.ID,
		Weight: pokemon.Weight,
		Height: pokemon.Height,
		Price:  pokemon.Price,
		Stock:  pokemon.Stock,
	}

	return c.JSON(http.StatusOK, output)
}
