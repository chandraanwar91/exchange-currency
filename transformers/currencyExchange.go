package transformers

import (
	"exchange-currency/models"
	"time"
)

type (
	CurrencyExchange struct {
		ID         int    `json:"id" gorm:"type:varchar(255)"`
		From       string `json:"from" gorm:"type:varchar(255)"`
		To         string `json:"to" gorm:"type:varchar(255)"`
		Status     string `json:"status" gorm:"type:enum('enable','disable','removed')"`
		CreatedBy  int    `json:"created_by" gorm:"column:created_by"`
		ModifiedBy int    `json:"modified_by" gorm:"column:modified_by"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	}
)

func (res *Transformer) Transform(currency *models.CurrencyExchange) *Transformer {
	res.Data = assignCurrencyExchange(currency)
	return res
}

func (res *CollectionTransformer) TransformCollection(currencies []*models.CurrencyExchange, pagination *models.Pagination) {
	for _, currency := range currencies {
		res.Data = append(res.Data, assignCurrencyExchange(currency))
	}

	res.Meta = models.Meta{Pagination: pagination}
}

func assignCurrencyExchange(currency *models.CurrencyExchange) interface{} {
	result := CurrencyExchange{}
	result.ID = currency.ID
	result.From = currency.From
	result.To = currency.To
	result.Status = currency.Status
	//result.CreatedBy = currency.CreatedBy
	//result.ModifiedBy = rate.ModifiedBy
	result.CreatedAt = currency.CreatedAt.Format(time.RFC3339)
	return result
}
