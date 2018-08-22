package DailyCurrencyExchange

import (
	"exchange-currency/models"
)

type IDailyCurrencyExchangeRepository interface {
	Store(m *models.DailyCurrencyExchange) (*models.DailyCurrencyExchange, error)
	FindAvailable(m *models.DailyCurrencyExchange) (bool,*models.DailyCurrencyExchange, error)
	Update(data *models.DailyCurrencyExchange,r *models.DailyCurrencyExchange) (*models.DailyCurrencyExchange,error)
	FetchByDate(page int, perPage int,date string) ([]*models.ListDailyCurrency,int, error)
	FetchTrendCurrency(CurrencyExchangeId int) ([]*models.ListTrendDailyCurrency, error)
}
