package handler

import (
	DCEInterfaces "exchange-currency/app/daily-currency-exchange"
	"exchange-currency/helpers"
	"exchange-currency/models"
	"strconv"
	"time"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"exchange-currency/transformers"
)

type createDailyCurrencyExchangeRules struct {
	Date string `valid:"required~parameter is empty"`
	From string `valid:"required~parameter is empty"`
	To   string `valid:"required~parameter is empty"`
	Rate string `valid:"required~parameter is empty"`

}

type DailyCurrencyExchangeHandler struct {
	DCEUsecase DCEInterfaces.IDailyCurrencyExchangeUseCase
}

func (a *DailyCurrencyExchangeHandler) CreateDailyCurrencyExchange(c *gin.Context) {
	var currencyExchange models.CurrencyExchange
	c.Bind(&currencyExchange)
	validation := ValidateDailyCurrency(c)
	if validation != nil {
		RespondFailValidation(c, validation)
		return
	}

	currencyId,err := a.DCEUsecase.GetCurrencyExchangeID(c,&currencyExchange)
	if err != nil {
		RespondFailValidation(c, "Failed create daily currency exchange Detail : "+err.Error())
		return
	}

	dailyCurrency := mapData(c,currencyId)

	_, errInsert := a.DCEUsecase.Store(c, dailyCurrency)
	if errInsert != nil {
		RespondFailValidation(c, "Failed create daily currency exchange Detail : "+errInsert.Error())
		return
	}
	RespondCreated(c, "Resource Created")
	return
}

func ValidateDailyCurrency(params *gin.Context) interface{} {
	DCERules := &createDailyCurrencyExchangeRules{
		From: params.PostForm("from"),
		To:   params.PostForm("to"),
		Date: params.PostForm("date"),
		Rate: params.PostForm("rate"),
	}
	_, err := govalidator.ValidateStruct(DCERules)
	if err != nil {
		respErr := helpers.ValidationError(DCERules, err)

		if respErr != nil {
			return respErr
		}
	}

	return nil
}

func mapData(params *gin.Context,CurrencyExchangeId int) *models.DailyCurrencyExchange{
	exchangeRate, _ := strconv.ParseFloat(params.PostForm("rate"), 64)
	createdBy, _ := strconv.Atoi(params.PostForm("created_by"))
	currencyDate,_ := time.Parse("2006-01-02", params.PostForm("date"))
	dailyCurrency := new(models.DailyCurrencyExchange)
	dailyCurrency.CurrencyDate 			= currencyDate
	dailyCurrency.ExchangeRate			= exchangeRate
	dailyCurrency.CurrencyExchangeId 	= CurrencyExchangeId
	dailyCurrency.CreatedBy				= createdBy
	return dailyCurrency
}

func (a *DailyCurrencyExchangeHandler) GetDailyExchangeByDate(c *gin.Context) {
	date := c.Query("date")
	dailyCur, pagination, err := a.DCEUsecase.FetchByDate(c,date)
	if err != nil {
		RespondFailValidation(c,err.Error())
		return
	}

	//Transform
	res := new(transformers.CollectionTransformer)
	res.TransformDailyCurrencyListCollection(dailyCur, pagination)
	RespondJSON(c,res)
	return
}

func (a *DailyCurrencyExchangeHandler) GetTrendCurrency(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	trendCurrencies, variance, err := a.DCEUsecase.FetchTrendCurrency(c,from,to)
	if err != nil {
		RespondFailValidation(c,err.Error())
		return
	}

	//Transform  
	res := new(transformers.CollectionTrendTransformer)
	res.TransformTrendCurrencyListCollection(trendCurrencies,variance,from,to)
	RespondJSON(c,res)
	return
}