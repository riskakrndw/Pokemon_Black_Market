package databases

import (
	"project/pbm/config"
	"project/pbm/models"
)

func CountTransaction() (int, error) {
	var count int
	if err := config.DB.Raw("SELECT COUNT(transactions.id) FROM transactions").Scan(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

func GetTotalById(transaction_id uint) (int, error) {
	var transaction models.Transaction
	status := "Success"
	if err := config.DB.Model(&transaction).Select("total").Where("id = ? AND status = ?", transaction_id, status).First(&transaction.Total).Error; err != nil {
		return transaction.Total, err
	}

	return transaction.Total, nil
}

func GetTotalSuccess() (int, error) {
	var transaction models.Transaction
	status := "Success"
	if err := config.DB.Model(&transaction).Select("COALESCE(sum(total),0)").Where("status = ?", status).First(&transaction.Total).Error; err != nil {
		return transaction.Total, err
	}

	return transaction.Total, nil
}

func GetTotalCancel() (int, error) {
	var transaction models.Transaction
	status := "Cancelled"
	if err := config.DB.Model(&transaction).Select("COALESCE(sum(total),0)").Where("status = ?", status).First(&transaction.Total).Error; err != nil {
		return transaction.Total, err
	}

	return transaction.Total, nil
}

func CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	if err := config.DB.Save(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func CreateDetailTransaction(detail_transaction models.DetailTransaction, transaction_id, pokemon_id uint, price int) (models.DetailTransaction, error) {
	detail_transaction = models.DetailTransaction{
		Price:         price,
		Quantity:      detail_transaction.Quantity,
		TransactionID: transaction_id,
		PokemonID:     pokemon_id,
	}
	if err := config.DB.Save(&detail_transaction).Error; err != nil {
		return detail_transaction, err
	}

	UpdateTotalTransaction(detail_transaction)
	UpdateStock(detail_transaction)

	return detail_transaction, nil
}

func UpdateTotalTransaction(detail_transaction models.DetailTransaction) (models.Transaction, error) {
	var transaction models.Transaction
	if err := config.DB.Find(&transaction, "id = ?", detail_transaction.TransactionID).Error; err != nil {
		return transaction, err
	}
	transaction.Total = transaction.Total + (detail_transaction.Price * detail_transaction.Quantity)
	if err := config.DB.Save(transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func UpdateStock(detail_transaction models.DetailTransaction) (models.Pokemon, error) {
	var pokemon models.Pokemon
	if err := config.DB.Find(&pokemon, "id = ?", detail_transaction.PokemonID).Error; err != nil {
		return pokemon, err
	}
	pokemon.Stock = pokemon.Stock - detail_transaction.Quantity
	if err := config.DB.Save(pokemon).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func CancelTransaction(transaction models.Transaction) (models.Transaction, error) {
	if err := config.DB.Save(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func UpdateStockAfterCancel(pokemon_id uint, quantity int) (models.Pokemon, error) {
	var pokemon models.Pokemon
	if err := config.DB.Find(&pokemon, "id = ?", pokemon_id).Error; err != nil {
		return pokemon, err
	}
	pokemon.Stock = pokemon.Stock + quantity
	if err := config.DB.Save(pokemon).Error; err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func GetAllTransactionSuccess() ([]models.Transaction, error) {
	var all_transaction []models.Transaction
	status := "Success"
	if err := config.DB.Find(&all_transaction, "status = ?", status).Error; err != nil {
		return all_transaction, err
	}
	return all_transaction, nil
}

func GetAllTransactionFailed() ([]models.Transaction, error) {
	var all_transaction []models.Transaction
	status := "Cancelled"
	if err := config.DB.Find(&all_transaction, "status = ?", status).Error; err != nil {
		return all_transaction, err
	}
	return all_transaction, nil
}

func GetTransactionById(id int) (models.Transaction, error) {
	var transaction models.Transaction

	if err := config.DB.Where("id = ?", id).First(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func GetDetailTransactionById(id int) ([]models.DetailTransaction, error) {
	var detail_transaction []models.DetailTransaction

	if err := config.DB.Find(&detail_transaction, "transaction_id = ?", id).Error; err != nil {
		return detail_transaction, err
	}

	return detail_transaction, nil
}
