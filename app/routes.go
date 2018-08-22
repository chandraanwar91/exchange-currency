package app

import (
	CEInterfaces "exchange-currency/app/currency-exchange"
	handlerCE "exchange-currency/app/currency-exchange/handler"
	DCEInterfaces "exchange-currency/app/daily-currency-exchange"
	handlerDCE "exchange-currency/app/daily-currency-exchange/handler"
	"github.com/gin-gonic/gin"
)

func NewCEHttpHandler(r *gin.Engine, us CEInterfaces.ICurrencyExchangeUseCase) {
	handler := &handlerCE.CurrencyExchangeHandler{
		CEUsecase: us,
	}
	rate := r.Group("/forex")
	rate.POST("/currency-exchange", handler.CreateCurrencyExchange)
	rate.GET("/currency-exchange", handler.GetAllCurrencyExchange)
	rate.DELETE("/currency-exchange/:id", handler.DeleteCurrencyExchange)
}

func NewDCEHttpHandler(r *gin.Engine, us DCEInterfaces.IDailyCurrencyExchangeUseCase) {
	handler := &handlerDCE.DailyCurrencyExchangeHandler{
		DCEUsecase: us,
	}
	rate := r.Group("/forex")
	rate.POST("/daily-currency-exchange", handler.CreateDailyCurrencyExchange)
	rate.GET("/daily-currency-exchange/lists", handler.GetDailyExchangeByDate)
	rate.GET("/daily-currency-exchange/trends", handler.GetTrendCurrency)
}
