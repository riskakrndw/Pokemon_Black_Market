package controllers

import (
	"net/http"
	"project/pbm/databases"
	"project/pbm/middlewares"
	"project/pbm/models"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type OutputTransaction2 struct {
	Total       int
	Transaction []Transaction1
}

type Transaction struct {
	EmployeeName string
	ID           uint
	Code         string
	Date         time.Time
	Total        int
	Status       string
}

type Transaction1 struct {
	EmployeeID uint
	ID         uint
	Code       string
	Date       time.Time
	Total      int
	Status     string
}

type OutputTransaction struct {
	EmployeeName string
	ID           uint
	Code         string
	Date         time.Time
	Total        int
	Status       string
	Pokemon      []Pokemon
}

type OutputTransaction1 struct {
	EmployeeID uint
	ID         uint
	Code       string
	Date       time.Time
	Total      int
	Status     string
	Pokemon    []Pokemon
}

type Pokemon struct {
	PokemonName string
	Quantity    int
	Price       int
}

func CreateTransaction(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID != 3 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Pengedar")
	}

	//get user's input
	input_user := []models.DetailTransaction{}
	c.Bind(&input_user)

	//check is data nil?
	for i := 0; i < len(input_user); i++ {
		if input_user[i].Quantity == 0 || input_user[i].PokemonID == 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Please fill all data",
			})
		}
	}

	//check is pokemon exist
	for i := 0; i < len(input_user); i++ {
		pokemon, err := databases.GetPokemonById(int(input_user[i].PokemonID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot get data")
		}
		if pokemon.ID == 0 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Pokemon not exists",
			})
		}
	}

	//check stock
	for i := 0; i < len(input_user); i++ {
		pokemon, err := databases.GetPokemonById(int(input_user[i].PokemonID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot get data")
		}
		if pokemon.Stock < input_user[i].Quantity {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Stock kurang",
			})
		}
	}

	//count transaction data for custom code transaction
	count_transaction, _ := databases.CountTransaction()
	code := strconv.Itoa(count_transaction + 1)

	//set transactions data
	set_transaction := models.Transaction{
		Code:        "PBM-" + code,
		DateOfMonth: time.Now(),
		Total:       0,
		Status:      "Success",
		UserID:      user.ID,
	}

	//create new transaction
	transaction, err := databases.CreateTransaction(set_transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot insert data")
	}

	//create new detail transactions
	var pokemon models.Pokemon
	var detail_transactions []models.DetailTransaction
	for i := 0; i < len(input_user); i++ {
		pokemon, err = databases.GetPokemonById(int(input_user[i].PokemonID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot get data")
		}
		detail_transaction, err := databases.CreateDetailTransaction(input_user[i], transaction.ID, pokemon.ID, pokemon.Price)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot insert data")
		}
		detail_transactions = append(detail_transactions, detail_transaction)
	}

	//get new total
	total, err := databases.GetTotalById(transaction.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//customize output
	var output_pokemons []Pokemon
	for i := 0; i < len(detail_transactions); i++ {
		pokemon, err = databases.GetPokemonById(int(detail_transactions[i].PokemonID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot get data")
		}
		new_result := Pokemon{
			PokemonName: pokemon.Name,
			Quantity:    detail_transactions[i].Quantity,
			Price:       detail_transactions[i].Price,
		}
		output_pokemons = append(output_pokemons, new_result)
	}
	output := OutputTransaction{
		EmployeeName: user.Name,
		ID:           transaction.ID,
		Code:         transaction.Code,
		Date:         transaction.DateOfMonth,
		Total:        total,
		Status:       transaction.Status,
		Pokemon:      output_pokemons,
	}

	return c.JSON(http.StatusOK, output)
}

func CreateTransactionTesting() echo.HandlerFunc {
	return CreateTransaction
}

func GetAllTransactionSuccess(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID != 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Bos")
	}

	//get total transaction
	total, err := databases.GetTotalSuccess()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//get all transactions
	all_transactions, _ := databases.GetAllTransactionSuccess()
	if len(all_transactions) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"Total":   total,
			"Message": "cannot find transaction",
		})
	}

	//customize output
	var transaction []Transaction1
	for i := 0; i < len(all_transactions); i++ {
		new_result := Transaction1{
			EmployeeID: all_transactions[i].UserID,
			ID:         all_transactions[i].ID,
			Code:       all_transactions[i].Code,
			Date:       all_transactions[i].DateOfMonth,
			Total:      all_transactions[i].Total,
			Status:     all_transactions[i].Status,
		}
		transaction = append(transaction, new_result)
	}
	output := OutputTransaction2{
		Total:       total,
		Transaction: transaction,
	}

	return c.JSON(http.StatusOK, output)
}

func GetAllTransactionSuccessTesting() echo.HandlerFunc {
	return GetAllTransactionSuccess
}

func GetAllTransactionFailed(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID != 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Bos")
	}

	//get total transaction
	total, err := databases.GetTotalCancel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//get all transactions
	all_transactions, _ := databases.GetAllTransactionFailed()
	if len(all_transactions) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"Total":   total,
			"Message": "cannot find transaction",
		})
	}

	//customize output
	var transaction []Transaction1
	for i := 0; i < len(all_transactions); i++ {
		new_result := Transaction1{
			EmployeeID: all_transactions[i].UserID,
			ID:         all_transactions[i].ID,
			Code:       all_transactions[i].Code,
			Date:       all_transactions[i].DateOfMonth,
			Total:      all_transactions[i].Total,
			Status:     all_transactions[i].Status,
		}
		transaction = append(transaction, new_result)
	}
	output := OutputTransaction2{
		Total:       total,
		Transaction: transaction,
	}

	return c.JSON(http.StatusOK, output)
}

func GetAllTransactionFailedTesting() echo.HandlerFunc {
	return GetAllTransactionFailed
}

func GetTransaction(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID != 1 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Bos")
	}

	//get param
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	//get transaction
	transaction, err := databases.GetTransactionById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//get detail transaction
	detail_transactions, err := databases.GetDetailTransactionById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//customize output
	var output_pokemons []Pokemon
	for i := 0; i < len(detail_transactions); i++ {
		pokemon, err := databases.GetPokemonById(int(detail_transactions[i].PokemonID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot get data")
		}
		new_result := Pokemon{
			PokemonName: pokemon.Name,
			Quantity:    detail_transactions[i].Quantity,
			Price:       detail_transactions[i].Price,
		}
		output_pokemons = append(output_pokemons, new_result)
	}
	output := OutputTransaction1{
		EmployeeID: transaction.UserID,
		ID:         transaction.ID,
		Code:       transaction.Code,
		Date:       transaction.DateOfMonth,
		Total:      transaction.Total,
		Status:     transaction.Status,
		Pokemon:    output_pokemons,
	}

	return c.JSON(http.StatusOK, output)
}

func GetTransactionTesting() echo.HandlerFunc {
	return GetTransaction
}

func CancelTransaction(c echo.Context) error {
	//check otorisasi
	logged_in_user_id := middlewares.ExtractToken(c)
	if logged_in_user_id == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please login first")
	}
	user, err := databases.GetUserById(logged_in_user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}
	if user.LevelID != 3 {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorized only by Pengedar")
	}

	//get param
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	//get transaction
	old_transaction, err := databases.GetTransactionById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//check expired
	today := time.Now()
	expired := old_transaction.DateOfMonth.Add(72 * time.Hour)
	if today.After(expired) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Garansi expired",
		})
	}

	//update data
	old_transaction.Status = "Cancelled"
	c.Bind(&old_transaction)
	new_transaction, err := databases.CancelTransaction(old_transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot update")
	}

	//get detail transaction
	detail_transactions, err := databases.GetDetailTransactionById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get data")
	}

	//update stock pokemon
	for i := 0; i < len(detail_transactions); i++ {
		databases.UpdateStockAfterCancel(detail_transactions[i].PokemonID, detail_transactions[i].Quantity)
	}

	//customize output
	var output_pokemons []Pokemon
	for i := 0; i < len(detail_transactions); i++ {
		pokemon, err := databases.GetPokemonById(int(detail_transactions[i].PokemonID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot get data")
		}
		new_result := Pokemon{
			PokemonName: pokemon.Name,
			Quantity:    detail_transactions[i].Quantity,
			Price:       detail_transactions[i].Price,
		}
		output_pokemons = append(output_pokemons, new_result)
	}
	output := OutputTransaction{
		EmployeeName: user.Name,
		ID:           new_transaction.ID,
		Code:         new_transaction.Code,
		Date:         new_transaction.DateOfMonth,
		Total:        new_transaction.Total,
		Status:       new_transaction.Status,
		Pokemon:      output_pokemons,
	}

	return c.JSON(http.StatusOK, output)
}

func CancelTransactionTesting() echo.HandlerFunc {
	return CancelTransaction
}
