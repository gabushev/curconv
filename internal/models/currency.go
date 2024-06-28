package models

import (
    "gorm.io/gorm"
)

type CurrencyPair struct {
	ID        uint   `gorm:"primaryKey"`
	From      string `gorm:"not null"`
	To        string `gorm:"not null"`
	Available bool   `gorm:"default:true"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&CurrencyPair{})
}
