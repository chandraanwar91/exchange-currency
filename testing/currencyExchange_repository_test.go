package testing

import (
	"testing"
	"strconv"
	CERepo "exchange-currency/app/currency-exchange/repository"
	"github.com/stretchr/testify/assert"
	"exchange-currency/db"
	"exchange-currency/models"
)


func TestStore(t *testing.T) {
	m := buildTestData()
	db := gorm.MysqlConn()
	assert := assert.New(t)
	//var err error
	a := CERepo.NewCERepository(db)
	currency,err := a.Store(m)

	assert.NoError(err)
	assert.Equal(int(1), currency.ID)
}

func TestFindAvailable(t *testing.T) {
	m := buildTestData()
	db := gorm.MysqlConn()
	assert := assert.New(t)
	a := CERepo.NewCERepository(db)
	status ,err := a.FindAvailable(m)
	assert.NoError(err)
	assert.True(status)
}

func TestFindByCurrency(t *testing.T) {
	m := buildTestData()
	db := gorm.MysqlConn()
	assert := assert.New(t)
	a := CERepo.NewCERepository(db)
	id ,err := a.FindByCurrency(m)
	assert.NoError(err)
	assert.Equal(id,1)
}

func TestGetById(t *testing.T) {
	db := gorm.MysqlConn()
	assert := assert.New(t)
	a := CERepo.NewCERepository(db)
	currency,err := a.GetById(1)
	assert.NoError(err)
	assert.Equal(1,currency.ID)
}

func TestFetch(t *testing.T) {
	m := buildTestData()
	db := gorm.MysqlConn()
	assert := assert.New(t)
	a := CERepo.NewCERepository(db)
	list, _ ,err := a.Fetch(1,1,m.Status)
	assert.NoError(err)
	assert.Len(list, 1)
}

func buildTestData() *models.CurrencyExchange{
	createdBy, _ := strconv.Atoi("1")
	currency := new(models.CurrencyExchange)
	currency.From		= "USD"
	currency.To			= "IDR"
	currency.Status		= "enable"
	currency.CreatedBy	= createdBy
	return currency
}