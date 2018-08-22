package handler

import (
	CEInterfaces "exchange-currency/app/currency-exchange"
	"exchange-currency/helpers"
	"exchange-currency/models"
	"exchange-currency/transformers"
	"strconv"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type createCurrencyExchangeRules struct {
	From string `valid:"required~parameter is empty"`
	To   string `valid:"required~parameter is empty"`
}

type CurrencyExchangeHandler struct {
	CEUsecase CEInterfaces.ICurrencyExchangeUseCase
}

func (a *CurrencyExchangeHandler) CreateCurrencyExchange(c *gin.Context) {

	validation := ValidateCurrency(c)
	if validation != nil {
		RespondFailValidation(c, validation)
		return
	}

	var currency models.CurrencyExchange
	c.Bind(&currency)

	_, err := a.CEUsecase.Store(c, &currency)
	if err != nil {
		RespondFailValidation(c, "Failed create rate Detail : "+err.Error())
		return
	}
	RespondCreated(c, "Resource Created")
	return
}

func (a *CurrencyExchangeHandler) GetAllCurrencyExchange(c *gin.Context) {

	currencies, pagination, err := a.CEUsecase.Fetch(c)
	if err != nil {
		RespondFailValidation(c, err.Error())
	}

	//Transform
	res := new(transformers.CollectionTransformer)
	res.TransformCollection(currencies, pagination)
	RespondJSON(c, res)
	return
}

func (a *CurrencyExchangeHandler) DeleteCurrencyExchange(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := a.CEUsecase.Delete(c, id)

	if err != nil {
		RespondFailValidation(c, "Failed to delete currency exchange")
		return
	}

	RespondDeleted(c, "Success Delete currency exchange")
	return
}

func ValidateCurrency(params *gin.Context) interface{} {
	CERules := &createCurrencyExchangeRules{
		From: params.PostForm("from"),
		To:   params.PostForm("to"),
	}
	_, err := govalidator.ValidateStruct(CERules)
	if err != nil {
		respErr := helpers.ValidationError(CERules, err)

		if respErr != nil {
			return respErr
		}
	}

	return nil
}
