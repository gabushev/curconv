package migrations

import (
	"curconv/internal/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedCurrencyPairs(db *gorm.DB) {
	pairs := []models.CurrencyPair{
		{From: "USDT", To: "EUR", Available: true},
		{From: "USDT", To: "USD", Available: true},
		{From: "USDT", To: "CNY", Available: true},
		{From: "USDT", To: "USDC", Available: true},
		{From: "USDT", To: "ETH", Available: true},
		{From: "USDC", To: "EUR", Available: true},
		{From: "USDC", To: "USD", Available: true},
		{From: "USDC", To: "CNY", Available: true},
		{From: "USDC", To: "ETH", Available: true},
		{From: "ETH", To: "EUR", Available: true},
		{From: "ETH", To: "USD", Available: true},
		{From: "ETH", To: "CNY", Available: true},
	}

	for _, pair := range pairs {
		if err := db.FirstOrCreate(&pair, models.CurrencyPair{From: pair.From, To: pair.To}).Error; err != nil {
			log.Printf("failed to seed currency pair %s to %s: %v", pair.From, pair.To, err)
		}
	}
}
