package models

import gorm2 "github.com/jinzhu/gorm"

const STATUS_ENABLE = "enable"
const STATUS_DISABLE = "disable"
const STATUS_REMOVED = "removed"

type (
	CurrencyExchange struct {
		BaseModel
		From       string `json:"from" form:"from"`
		To         string `json:"to" form:"to"`
		Status     string `json:"status" form:"status"`
		CreatedBy  int    `json:"created_by" form:"created_by" gorm:"column:created_by"`
		ModifiedBy int    `json:"modified_by" form:"modified_by" gorm:"column:modified_by"`
	}
)

func (CurrencyExchange) TableName() string {
	return "currency_exchange"
}

func ScopeActive(db *gorm2.DB) *gorm2.DB {
	return db.Where("status = ?", "enable")
}
