package CurrencyExchange

import (
	"exchange-currency/models"
)

type ICurrencyExchangeRepository interface {
	Store(m *models.CurrencyExchange) (*models.CurrencyExchange, error)
	FindAvailable(m *models.CurrencyExchange) (bool, error)
	FindByCurrency(m *models.CurrencyExchange) (int,error)
	GetById(id uint64) (*models.CurrencyExchange, error)
	Fetch(page, perPage int, status string) ([]*models.CurrencyExchange, int, error)
	Delete(data *models.CurrencyExchange, ModifiedBy int) error
}
