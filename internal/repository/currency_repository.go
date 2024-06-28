package repository

import (
	"curconv/internal/models"
	"gorm.io/gorm"
)

const selectWithOrderOptionQry = `SELECT * FROM "currency_pairs" cp WHERE (cp.available = ? and cp.from = ? and cp.to = ?) OR (cp.available = ? and cp.from = ? and cp.to = ?)`

type CurrencyRepository struct {
	DB *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepository {
	return &CurrencyRepository{DB: db}
}

func (r *CurrencyRepository) GetAvailablePairs() ([]models.CurrencyPair, error) {
	var pairs []models.CurrencyPair
	result := r.DB.Where("available = ?", true).Find(&pairs)
	return pairs, result.Error
}

func (r *CurrencyRepository) FindAvailablePair(from string, to string) (models.CurrencyPair, error) {
	var pair models.CurrencyPair
	result := r.DB.Raw(selectWithOrderOptionQry, true, from, to, true, to, from).Scan(&pair)
	if result.Error != nil {
		return pair, result.Error
	}
	return pair, nil
}
