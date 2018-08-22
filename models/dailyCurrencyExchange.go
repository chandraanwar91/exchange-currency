package models

import (
	"time"
)

type (
	DailyCurrencyExchange struct {
		BaseModel
		CurrencyDate       	time.Time `json:"currency_date" gorm:"type:date";column:currency_date`
		CurrencyExchangeId 	int      `json:"currency_exchange_id"`
		CreatedBy  			int    `json:"created_by" form:"created_by" gorm:"column:created_by"`
		ModifiedBy int    `json:"modified_by" form:"modified_by" gorm:"column:modified_by"`
		ExchangeRate       	float64   `json:"exchange_rate" gorm:"type:decimal(12,2)"`
	}
	

	ListDailyCurrency struct {
		From       	string 		`json:"from" form:"from"`
		To         	string 		`json:"to" form:"to"`
		Date       	time.Time 	`json:"date" gorm:"type:date";column:currency_date`
		Rate 		float64     `json:"rate"`
		AvgDaily float64    	`json:"avg_daily"`
	}

	ListTrendDailyCurrency struct {
		MinRate	float64 	`json:"max_rate"`	
		MaxRate	float64     `json:"min_rate"`
		AvgRate float64    	`json:"avg_rate"`
		Date	time.Time 	`json:"date" gorm:"type:date";column:currency_date`	
		Rate	float64     `json:"rate"`
	}
	
)

func (DailyCurrencyExchange) TableName() string {
	return "daily_currency_exchange"
}