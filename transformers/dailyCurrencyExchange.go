package transformers

import (
	"exchange-currency/models"
	//"time"
	"strconv"
)

type (
	DailyCurrencyExchangeList struct {
		From	string 	`json:"from" gorm:"type:varchar(255)"`
		To		string 	`json:"to" gorm:"type:varchar(255)"`
		Rate	string	`json:"rate"`
		DayAvg	string	`json:"7_day_avg"`
	}

	CollectionTrendTransformer struct {
		Data TrendCurrencyExchange `json:"data"`
	}

	TrendCurrencyExchange struct {
		From	string 	`json:"from" gorm:"type:varchar(255)"`
		To		string 	`json:"to" gorm:"type:varchar(255)"`
		Variance string `json:"variance" gorm:"type:varchar(255)"`
		Detail	[]DetailTrendCurrencyExchange	`json:"details"`
	}

	DetailTrendCurrencyExchange struct {
		Rate	string	`json:"rate"`
		Date	string	`json:"date"`
	}
)

func (res *Transformer) DailyCurrencyListTransform(currency *models.ListDailyCurrency) *Transformer {
	res.Data = assignDailyCurrencyExchangeList(currency)
	return res
}

func (res *CollectionTransformer) TransformDailyCurrencyListCollection(currencies []*models.ListDailyCurrency, pagination *models.Pagination) {
	for _, currency := range currencies {
		res.Data = append(res.Data, assignDailyCurrencyExchangeList(currency))
	}

	res.Meta = models.Meta{Pagination: pagination}
}

func (res *CollectionTrendTransformer) TransformTrendCurrencyListCollection(trendCurrencies []*models.ListTrendDailyCurrency, variance float64,from string,to string) {
	varianceStr := strconv.FormatFloat(variance, 'f', 6, 64)
	trendStruct := TrendCurrencyExchange{}
	for _, trendCur := range trendCurrencies {
		
		trendStruct.Detail = append(trendStruct.Detail, assignTrendCurrencyList(trendCur))
	}

	trendStruct.Variance = varianceStr
	trendStruct.From = from
	trendStruct.To =  to

	res.Data = trendStruct
}


func (res *CollectionTransformer) TransformTrendCurrencyCollection(currencies []*models.ListDailyCurrency, pagination *models.Pagination) {
	for _, currency := range currencies {
		res.Data = append(res.Data, assignDailyCurrencyExchangeList(currency))
	}

	res.Meta = models.Meta{Pagination: pagination}
}

func assignDailyCurrencyExchangeList(currency *models.ListDailyCurrency) interface{} {
	rate := strconv.FormatFloat(currency.Rate, 'f', 6, 64)
	avgDaily := strconv.FormatFloat(currency.AvgDaily, 'f', 6, 64)
	if currency.Rate <= 0{
		rate = "insufficient data"
	}

	if currency.AvgDaily <= 0{
		avgDaily = ""
	}

	result := DailyCurrencyExchangeList{}
	result.From = currency.From
	result.To = currency.To
	result.Rate = rate
	result.DayAvg = avgDaily
	return result
}

func assignTrendCurrencyList(trendCurrency *models.ListTrendDailyCurrency) DetailTrendCurrencyExchange {
	rate := strconv.FormatFloat(trendCurrency.Rate, 'f', 6, 64)
	date := trendCurrency.Date.Format("2006-01-02")
	
	result := DetailTrendCurrencyExchange{}
	result.Rate = rate
	result.Date = date
	return result
}
