package utils

import "gorm.io/gorm"

/*
* Struct for transaction as json form
 */

type Transaction struct {
	Id_transaction string             `json:"id_transaction" validate:"required"`
	Id_user        uint               `json:"user" validate:"required"`
	Item           []Transaction_Item `gorm:"foreignKey:id_transaction; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"item" validate:"required"`
	Type           string             `json:"type" validate:"required"`
	Method_pay     string             `json:"methode_pay" validate:"required" `
	Total          uint               `json:"total" validate:"required"`
	Amount         uint               `json:"amount" validate:"required"`
}

type Transaction_Item struct {
	Id_food    uint `json:"id_food" validate:"required"`
	Id_variant uint `json:"id_variant" validate:"required"`
	Price      uint `json:"price" validate:"required"`
	Qty        uint `json:"qty" validate:"required" `
}

/*
* Struct for get data  relasion using INNER JOIN
 */

type Result struct {
	Id          string `json:"id"`
	Fullname    string `json:"fullname"`
	Price       int    `json:"price"`
	Pay_methode string `json:"pay_methode"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Food_Name   string `json:"food_name"`
	Variant     string `json:"variant"`
	Qty         int    `json:"qty"`
}

type Variant struct {
	Id        string         `json:"id"`
	Variant   string         `json:"variant"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
