package repository

import (
	"errors"
	"exchange-currency/models"
	CEInterface "exchange-currency/app/currency-exchange"
	"time"
	gorm1 "github.com/jinzhu/gorm"
)

type CERepository struct {
	Conn *gorm1.DB
}

func NewCERepository(Conn *gorm1.DB) CEInterface.ICurrencyExchangeRepository {

	return &CERepository{Conn}
}

func (m *CERepository) Store(a *models.CurrencyExchange) (*models.CurrencyExchange, error) {
	var err error

	a.CreatedAt = time.Now()
	tx := m.Conn.Begin()
	if err = tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return a, nil
}

func (m *CERepository) FindAvailable(a *models.CurrencyExchange) (bool, error) {
	var (
		err   error
		total int
	)

	//Initialization
	tx := m.Conn.Begin()

	tx = tx.Table("currency_exchange").Where("`from` = ?", a.From).Where("`to` = ?", a.To)

	if err = tx.Count(&total).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	if total > 0 {
		return true, nil
	}

	return false, nil
}

func (m *CERepository) Fetch(page, perPage int, status string) ([]*models.CurrencyExchange, int, error) {
	var (
		currencies []*models.CurrencyExchange
		total      int
		err        error
	)

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 10
	}

	offset := (page * perPage) - perPage
	//Initialization
	tx := m.Conn.Begin()

	txCount := tx

	//Count All
	//txCount.Table("rate_backoffice").Where("applicable_to = ? ",ApplicableTo).Scopes(models.ScopeActive).Count(&total)
	txCount = txCount.Table("currency_exchange")

	if status != "" {
		txCount = txCount.Where("status = ?", status)
		tx = tx.Where("status = ? ", status)
	}

	txCount = txCount.Scopes(models.ScopeActive)

	txCount.Count(&total)
	err = tx.Scopes(models.ScopeActive).Limit(perPage).Offset(offset).Find(&currencies).Error
	if err != nil {
		tx.Rollback()
		return currencies, total, err
	}
	tx.Commit()

	return currencies, total, err
}

func (m *CERepository) GetById(id uint64) (*models.CurrencyExchange, error) {
	var (
		currency models.CurrencyExchange
		err      error
	)

	//Initialization
	tx := m.Conn.Begin()
	err = tx.First(&currency, id).Error
	if err != nil {
		tx.Rollback()
		return &currency, nil
	}
	tx.Commit()

	if currency.ID == 0 {
		return nil, errors.New("ID not found")
	}

	return &currency, nil
}

func (m *CERepository) Delete(a *models.CurrencyExchange, ModifiedBy int) error {
	var err error

	a.UpdatedAt = time.Now()
	a.Status = "removed"
	a.ModifiedBy = ModifiedBy
	tx := m.Conn.Begin()
	if err = tx.Save(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (m *CERepository) FindByCurrency(a *models.CurrencyExchange) (int, error) {
	var (
		err   error
		currency models.CurrencyExchange
	)

	//Initialization
	tx := m.Conn.Begin()

	tx = tx.Table("currency_exchange").Where("`from` = ?", a.From).Where("`to` = ?", a.To)

	if err = tx.First(&currency).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	if currency.ID > 0 {
		return currency.ID, nil
	}

	return 0, nil
}