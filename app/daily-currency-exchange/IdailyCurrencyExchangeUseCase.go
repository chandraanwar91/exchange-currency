package DailyCurrencyExchange

import (
	"exchange-currency/models"

	"github.com/gin-gonic/gin"
)

type IDailyCurrencyExchangeUseCase interface {
	Store(c *gin.Context, dce *models.DailyCurrencyExchange) (*models.DailyCurrencyExchange, error)
	GetCurrencyExchangeID(c *gin.Context,m *models.CurrencyExchange) (int, error)
	FetchByDate(c *gin.Context,date string) ([]*models.ListDailyCurrency, *models.Pagination, error)
	FetchTrendCurrency(c *gin.Context,from string,to string) ([]*models.ListTrendDailyCurrency,float64, error)
}
