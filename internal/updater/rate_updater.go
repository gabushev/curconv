package updater

import (
	"curconv/internal/repository"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type RateGetter interface {
	GetPrices(pairs []string) (map[string]float64, error)
}

type RateUpdater struct {
	CurrencyRepo *repository.CurrencyRepository
	Rates        map[string]float64
	Mutex        sync.RWMutex
	rg           RateGetter
}

func NewRateUpdater(repo *repository.CurrencyRepository, rg RateGetter) *RateUpdater {
	return &RateUpdater{
		CurrencyRepo: repo,
		Rates:        make(map[string]float64),
		Mutex:        sync.RWMutex{},
		rg:           rg,
	}
}

func (u *RateUpdater) Start() {
	go func() {
		for {
			u.updateRates()
			time.Sleep(1 * time.Minute)
		}
	}()
}

func (u *RateUpdater) updateRates() {
	log.Info().Msg("Updating rates")
	pairs, err := u.CurrencyRepo.GetAvailablePairs()
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch currency pairs from database")
		return
	}

	log.Info().Msg(fmt.Sprintf("Updating rates, got pairs: %d", len(pairs)))
	var pairStrings []string
	for _, pair := range pairs {
		// Add both directions of the pair to reduce chance of human mistake on the data setting
		pairStrings = append(pairStrings, fmt.Sprintf("%s/%s", pair.From, pair.To))
		pairStrings = append(pairStrings, fmt.Sprintf("%s/%s", pair.To, pair.From))
	}

	rates, err := u.rg.GetPrices(pairStrings)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch rates from forex API")
		return
	}
	for i := range rates {
		currenciesKeyRev := strings.Split(i, "/")
		reversedKey := fmt.Sprintf("%s/%s", currenciesKeyRev[1], currenciesKeyRev[0])
		if _, ok := rates[reversedKey]; !ok {
			rates[reversedKey] = 1 / rates[i]
		}
	}
	u.Mutex.Lock()
	u.Rates = rates
	u.Mutex.Unlock()
}

func (u *RateUpdater) GetRate(from, to string) (float64, bool) {
	_, err := u.CurrencyRepo.FindAvailablePair(from, to)
	if err != nil {
		return 0, false
	}

	u.Mutex.RLock()
	defer u.Mutex.RUnlock()
	rate, exists := u.Rates[fmt.Sprintf("%s/%s", from, to)]
	return rate, exists
}
