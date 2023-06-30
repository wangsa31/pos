package models

import (
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id          int    `gorm:"primarykey;autoIncrement"`
	Email       string `gorm:"type:varchar(100);not null" json:"email" validate:"required,email"`
	Password    string `gorm:"type:varchar(255);not null" json:"password" validate:"required"`
	Fullname    string `gorm:"type:varchar(100);not null" json:"fullname" validate:"required"`
	Gender      string `gorm:"not null;type:enum('male','female')" json:"gender" validate:"required"`
	Address     string `gorm:"type:varchar(100);not null" json:"address" validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Sift        []Sift         `gorm:"foreignKey:id_user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transaction []Transaction  `gorm:"foreignKey:id_user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Sift struct {
	Id         int  `gorm:"primarykey;autoIncrement"`
	Id_user    uint `gorm:"type: int(255); not null;"`
	Start_sift time.Time
	End_sift   time.Time
	Cash_open  uint `gorm:"type: int(255); not null"`
	Cash_close uint `gorm:"type: int(255); not null"`
}

type Food struct {
	Id               int    `gorm:"primarykey;autoIncrement"`
	Food_Name        string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Id_Variant       int    `gorm:"type:int(255);not null" json:"variant" validate:"required"`
	Discount         int    `gorm:"type:int(255);not null;default:0" json:"discount,omitempty"`
	Price            int    `gorm:"type:int(255);not null" json:"price" validate:"required"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt   `gorm:"index"`
	Item_transaction Item_transaction `gorm:"foreignKey:id_foods;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Variant struct {
	Id               int                `gorm:"primarykey;autoIncrement" json:"id"`
	Variant          string             `gorm:"type:varchar(100);not null" json:"variant"`
	CreatedAt        time.Time          `json:"create_at"`
	UpdatedAt        time.Time          `json:"update_at"`
	DeletedAt        gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
	Food             []Food             `gorm:"foreignKey:id_variant;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Item_transaction []Item_transaction `gorm:"foreignKey:id_variant;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Transaction struct {
	Id               string `gorm:"primarykey;type:varchar(255)"`
	Id_user          int    `gorm:"type:int(255);not null"`
	Total_amount     int    `gorm:"type:int(255);not null"`
	Pay_methode      string `gorm:"type:varchar(255);not null"`
	Status           string `gorm:"not null;type:enum('success','failed')"`
	Type             string `gorm:"not null;type:enum('dain in','take way')"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Item_transaction Item_transaction `gorm:"foreignKey:id_transaction;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Item_transaction struct {
	Id             int    `gorm:"primarykey;autoIncrement"`
	Id_transaction string `gorm:"type: varchar(255); not null"`
	Id_foods       int    `gorm:"type: int(255) UNSIGNED; not null;"`
	Id_variant     int    `gorm:"type: int(255) UNSIGNED; not null;"`
	Qty            int    `gorm:"type:int(255); not null"`
}

type Database struct {
	Db *gorm.DB
}

var (
	once     sync.Once
	instance *Database
)

func Getdatabase() *Database {
	once.Do(func() {
		dsn := "root:@tcp(127.0.0.1:3306)/pos?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
		}
		instance = &Database{
			Db: db,
		}

	})

	return instance
}
