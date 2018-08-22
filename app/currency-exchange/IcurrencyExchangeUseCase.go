package CurrencyExchange

import (
	"exchange-currency/models"
	"github.com/gin-gonic/gin"
)

type ICurrencyExchangeUseCase interface {
	Store(c *gin.Context, ce *models.CurrencyExchange) (*models.CurrencyExchange, error)
	Fetch(c *gin.Context) ([]*models.CurrencyExchange, *models.Pagination, error)
	Delete(c *gin.Context, id uint64) error
}
