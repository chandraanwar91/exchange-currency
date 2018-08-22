package testing

import (
	"testing"
	"time"
	"strconv"
	DCERepo "exchange-currency/app/daily-currency-exchange/repository"
	"github.com/stretchr/testify/assert"
	"exchange-currency/db"
	"exchange-currency/models"
)
func TestDailyStore(t *testing.T) {
	m := buildTestDailyCurData()
	db := gorm.MysqlConn()
	assert := assert.New(t)
	//var err error
	a := DCERepo.NewDCERepository(db)
	dailyCurrency,err := a.Store(m)

	assert.NoError(err)
	assert.Equal(int(1), dailyCurrency.ID)
}

func TestDailyFindAvailable(t *testing.T) {
	m := buildTestDailyCurData()
	db := gorm.MysqlConn()
	assert := assert.New(t)
	a := DCERepo.NewDCERepository(db)
	status,dailyCur,err := a.FindAvailable(m)
	assert.NoError(err)
	assert.True(status)
	assert.Equal(1,dailyCur.ID)
}

func TestDailyFindByDate(t *testing.T) {
	db := gorm.MysqlConn()
	assert := assert.New(t)
	a := DCERepo.NewDCERepository(db)
	dailyCur,total,err := a.FetchByDate(1,1,"2018-08-22")
	assert.NoError(err)
	assert.Len(dailyCur,1)
	assert.Equal(1,total)
}

func TestFetchTrendCurrency(t *testing.T) {
	m := buildTestDailyCurData()
	db := gorm.MysqlConn()
	assert := assert.New(t)
	a := DCERepo.NewDCERepository(db)
	trendCur,err := a.FetchTrendCurrency(m.ID)
	assert.NoError(err)
	assert.Len(trendCur,1)
}

func buildTestDailyCurData() *models.DailyCurrencyExchange{
	createdBy, _ := strconv.Atoi("1")
	currencyDate,_ := time.Parse("2006-01-02", "2018-08-22")
	rate, _ := strconv.ParseFloat("0.7570900", 64)
	dailyCurrency := new(models.DailyCurrencyExchange)
	dailyCurrency.CurrencyDate = currencyDate
	dailyCurrency.CurrencyExchangeId	= 1
	dailyCurrency.ExchangeRate = rate
	dailyCurrency.CreatedBy	= createdBy
	return dailyCurrency
}