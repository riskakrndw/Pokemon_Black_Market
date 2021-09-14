package models

type DetailTransaction struct {
	Price    int `json:"price" form:"price"`
	Quantity int `json:"quantity" form:"quantity"`

	//PK
	TransactionID uint `gorm:"primaryKey" json:"transaction_id" form:"transaction_id"`
	PokemonID     uint `gorm:"primaryKey" json:"pokemon_id" form:"pokemon_id"`
}
