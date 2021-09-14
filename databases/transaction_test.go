package databases

import (
	"project/pbm/config"
	"project/pbm/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ()

func TestCreateTransactionSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Success",
		UserID:      mock_user.ID,
	}
	transaction, err := CreateTransaction(mock_transaction)
	if assert.NoError(t, err) {
		assert.Equal(t, "PBM-1", transaction.Code)
		assert.Equal(t, 10000, transaction.Total)
		assert.Equal(t, "Success", transaction.Status)
	}
}

func TestCreateTransactionError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Success",
		UserID:      mock_user.ID,
	}
	_, err := CreateTransaction(mock_transaction)
	assert.Error(t, err)
}

func TestCreateDetailTransactionSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.DetailTransaction{})
	config.DB.Migrator().AutoMigrate(&models.DetailTransaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	pokemon := models.Pokemon{
		Name:   "bulbasaur",
		Weight: 10,
		Height: 20,
		Price:  500000,
		Stock:  100,
	}
	if err := config.DB.Save(&pokemon).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Success",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	mock_detail_transaction := models.DetailTransaction{
		Price:    pokemon.Price,
		Quantity: 10,
	}
	detail_transaction, err := CreateDetailTransaction(mock_detail_transaction, mock_transaction.ID, pokemon.ID, pokemon.Price)
	if assert.NoError(t, err) {
		assert.Equal(t, 500000, detail_transaction.Price)
		assert.Equal(t, 10, detail_transaction.Quantity)
		assert.Equal(t, uint(1), detail_transaction.PokemonID)
	}
}

func TestCreateDetailTransactionError(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.DetailTransaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	pokemon := models.Pokemon{
		Name:   "bulbasaur",
		Weight: 10,
		Height: 20,
		Price:  500000,
		Stock:  100,
	}
	if err := config.DB.Save(&pokemon).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Success",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	mock_detail_transaction := models.DetailTransaction{
		Price:    pokemon.Price,
		Quantity: 10,
	}
	_, err := CreateDetailTransaction(mock_detail_transaction, mock_transaction.ID, pokemon.ID, pokemon.Price)
	assert.Error(t, err)
}

func TestGetAllTransactionSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Success",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	transaction, err := GetAllTransactionSuccess()
	if assert.NoError(t, err) {
		assert.Equal(t, "PBM-1", transaction[0].Code)
		assert.Equal(t, 10000, transaction[0].Total)
		assert.Equal(t, "Success", transaction[0].Status)
	}
}

func TestGetAllTransactionFailed(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Cancelled",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	transaction, err := GetAllTransactionFailed()
	if assert.NoError(t, err) {
		assert.Equal(t, "PBM-1", transaction[0].Code)
		assert.Equal(t, 10000, transaction[0].Total)
		assert.Equal(t, "Cancelled", transaction[0].Status)
	}
}

func TestGetTransactionById(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Success",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	transaction, err := GetTransactionById(int(mock_transaction.ID))
	if assert.NoError(t, err) {
		assert.Equal(t, "PBM-1", transaction.Code)
		assert.Equal(t, 10000, transaction.Total)
		assert.Equal(t, "Success", transaction.Status)
	}
}

func TestGetDetailTransactionById(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.DetailTransaction{})
	config.DB.Migrator().AutoMigrate(&models.DetailTransaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	pokemon := models.Pokemon{
		Name:   "bulbasaur",
		Weight: 10,
		Height: 20,
		Price:  500000,
		Stock:  100,
	}
	if err := config.DB.Save(&pokemon).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Success",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	mock_detail_transaction := models.DetailTransaction{
		Price:         pokemon.Price,
		Quantity:      10,
		TransactionID: mock_transaction.ID,
		PokemonID:     pokemon.ID,
	}
	if err := config.DB.Save(&mock_detail_transaction).Error; err != nil {
		t.Error(err)
	}
	detail_transaction, err := GetDetailTransactionById(int(mock_detail_transaction.TransactionID))
	if assert.NoError(t, err) {
		assert.Equal(t, 500000, detail_transaction[0].Price)
		assert.Equal(t, 10, detail_transaction[0].Quantity)
		assert.Equal(t, uint(1), detail_transaction[0].PokemonID)
	}
}

func TestCancelTransactionSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Cancelled",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	transaction, err := CancelTransaction(mock_transaction)
	if assert.NoError(t, err) {
		assert.Equal(t, "PBM-1", transaction.Code)
		assert.Equal(t, 10000, transaction.Total)
		assert.Equal(t, "Cancelled", transaction.Status)
	}
}

func TestUpdateStockAfterCancelSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	config.DB.Migrator().DropTable(&models.Pokemon{})
	config.DB.Migrator().AutoMigrate(&models.Pokemon{})
	config.DB.Migrator().DropTable(&models.DetailTransaction{})
	config.DB.Migrator().AutoMigrate(&models.DetailTransaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	pokemon := models.Pokemon{
		Name:   "bulbasaur",
		Weight: 10,
		Height: 20,
		Price:  500000,
		Stock:  100,
	}
	if err := config.DB.Save(&pokemon).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Success",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	mock_detail_transaction := models.DetailTransaction{
		Price:         pokemon.Price,
		Quantity:      10,
		TransactionID: mock_transaction.ID,
		PokemonID:     pokemon.ID,
	}
	if err := config.DB.Save(&mock_detail_transaction).Error; err != nil {
		t.Error(err)
	}
	pokemon, err := UpdateStockAfterCancel(pokemon.ID, mock_detail_transaction.Quantity)
	if assert.NoError(t, err) {
		assert.Equal(t, "bulbasaur", pokemon.Name)
		assert.Equal(t, 10, pokemon.Weight)
		assert.Equal(t, 20, pokemon.Height)
		assert.Equal(t, 110, pokemon.Stock)
	}
}

func TestCountTransactionSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Success",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	transaction, err := CountTransaction()
	if assert.NoError(t, err) {
		assert.Equal(t, 1, transaction)
	}
}

func TestGetTotalByIdSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Success",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	transaction, err := GetTotalById(mock_transaction.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, 10000, transaction)
	}
}

func TestGetTotalSuccess(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Success",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	transaction, err := GetTotalSuccess()
	if assert.NoError(t, err) {
		assert.Equal(t, 10000, transaction)
	}
}

func TestGetTotalCancelled(t *testing.T) {
	config.ConfigTest()
	config.DB.Migrator().DropTable(&models.User{})
	config.DB.Migrator().AutoMigrate(&models.User{})
	config.DB.Migrator().DropTable(&models.Transaction{})
	config.DB.Migrator().AutoMigrate(&models.Transaction{})
	if err := config.DB.Save(&mock_user).Error; err != nil {
		t.Error(err)
	}
	mock_transaction := models.Transaction{
		Code:        "PBM-1",
		DateOfMonth: time.Now(),
		Total:       10000,
		Status:      "Cancelled",
		UserID:      mock_user.ID,
	}
	if err := config.DB.Save(&mock_transaction).Error; err != nil {
		t.Error(err)
	}
	transaction, err := GetTotalCancel()
	if assert.NoError(t, err) {
		assert.Equal(t, 10000, transaction)
	}
}
