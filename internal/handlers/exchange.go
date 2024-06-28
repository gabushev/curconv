package handlers

import (
	"curconv/internal/updater"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ExchangeHandler struct {
	RateUpdater *updater.RateUpdater
}

func NewExchangeHandler(rateUpdater *updater.RateUpdater) *ExchangeHandler {
	return &ExchangeHandler{RateUpdater: rateUpdater}
}

//		@Summary Get amount converted from one currency to another
//	 @Description Get the exchange rate between two currencies with a given amount
//	 @ID get-exchange-rate
//	 @Produce  json
//	 @Param   from     query    string     true        "From currency code"
//	 @Param   to       query    string     true        "To currency code"
//	 @Param   amount   query    float64    true        "Amount to convert"
//	 @Success 200 {object} map[string]interface{}
//	 @Failure 400 {object} map[string]string
//	 @Router /convert [get]
func (h *ExchangeHandler) Convert(c *fiber.Ctx) error {
	from := c.Query("from")
	to := c.Query("to")
	amount, err := strconv.ParseFloat(c.Query("amount"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid amount")
	}

	rate, exists := h.RateUpdater.GetRate(from, to)
	if !exists {
		return c.Status(fiber.StatusBadRequest).SendString("Exchange rate not available")
	}

	convertedAmount := amount * rate

	return c.JSON(fiber.Map{
		"from":   from,
		"to":     to,
		"amount": amount,
		"result": convertedAmount,
	})
}
