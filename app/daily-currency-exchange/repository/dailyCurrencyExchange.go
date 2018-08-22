package repository

import (
	"exchange-currency/models"
	DCEInterface "exchange-currency/app/daily-currency-exchange"
	"time"
	gorm1 "github.com/jinzhu/gorm"
)

type DCERepository struct {
	Conn *gorm1.DB
}

func NewDCERepository(Conn *gorm1.DB) DCEInterface.IDailyCurrencyExchangeRepository {

	return &DCERepository{Conn}
}

func (m *DCERepository) Store(a *models.DailyCurrencyExchange) (*models.DailyCurrencyExchange, error) {
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

func (m *DCERepository) FindAvailable(a *models.DailyCurrencyExchange) (bool,*models.DailyCurrencyExchange, error) {
	var (
		err   error
		dailyCur models.DailyCurrencyExchange
	)

	//Initialization
	tx := m.Conn.Begin()

	tx = tx.Table("daily_currency_exchange").Where("`currency_date` = ?", a.CurrencyDate).Where("`currency_exchange_id` = ?", a.CurrencyExchangeId)

	if err = tx.First(&dailyCur).Error; err != nil {
		tx.Rollback()
		return false,nil, err
	}

	tx.Commit()

	if dailyCur.ID > 0 {
		return true, &dailyCur,nil
	}

	return false,nil, nil
}

func (m *DCERepository) Update(data *models.DailyCurrencyExchange,r *models.DailyCurrencyExchange) (*models.DailyCurrencyExchange,error) {
	var err error

	r.UpdatedAt 	= time.Now()
	r.ExchangeRate = data.ExchangeRate
	r.ModifiedBy	= data.ModifiedBy
	tx := m.Conn.Begin()
	if err = tx.Save(&r).Error; err != nil {
		tx.Rollback()
		return r,err
	}
	tx.Commit()

	return r,err
}

func (m *DCERepository) FetchByDate(page, perPage int,date string) ([]*models.ListDailyCurrency,int, error) {
	var (
		err   error
		total      int
		dailyCur []*models.ListDailyCurrency
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

	if err = tx.Raw("SELECT `from`,`to`,daily_currency_exchange.`currency_date` AS `date`,daily_currency_exchange.`exchange_rate` AS rate,(SELECT AVG(exchange_rate) FROM daily_currency_exchange WHERE daily_currency_exchange.currency_exchange_id = currency_exchange.id AND currency_date BETWEEN DATE_SUB((?), INTERVAL 7 DAY) AND (?)) AS avg_daily FROM currency_exchange LEFT JOIN daily_currency_exchange ON currency_exchange.id = daily_currency_exchange.currency_exchange_id AND daily_currency_exchange.currency_date = (?) GROUP BY daily_currency_exchange.currency_exchange_id ORDER BY currency_exchange.id",date,date,date).Limit(perPage).Offset(offset).Scan(&dailyCur).Error; err != nil {
		tx.Rollback()
		return nil,total,err
	}

	tx.Commit()
	total = len(dailyCur)
	return dailyCur,total, nil
}

func (m *DCERepository) FetchTrendCurrency(CurrencyExchangeId int) ([]*models.ListTrendDailyCurrency, error) {
	var (
		err   		error
		trendDailyCur []*models.ListTrendDailyCurrency
	)

	//Initialization
	tx := m.Conn.Begin()

	if err = tx.Raw("SELECT MIN(exchange_rate) AS min_rate, MAX(exchange_rate) AS max_rate,currency_date as date,exchange_rate as rate FROM daily_currency_exchange  where currency_exchange_id = (?) ORDER BY daily_currency_exchange.currency_date desc LIMIT 7",CurrencyExchangeId).Scan(&trendDailyCur).Error; err != nil {
		tx.Rollback()
		return nil,err
	}

	tx.Commit()

	return trendDailyCur, nil
}
