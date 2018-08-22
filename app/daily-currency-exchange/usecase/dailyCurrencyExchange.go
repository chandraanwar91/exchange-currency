package usecase

import (
	"errors"
	DCEInterfaces "exchange-currency/app/daily-currency-exchange"
	CEInterfaces "exchange-currency/app/currency-exchange"
	"exchange-currency/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DCEUsecase struct {
	DCERepo DCEInterfaces.IDailyCurrencyExchangeRepository
	CERepo CEInterfaces.ICurrencyExchangeRepository
}

func NewDailyCurrencyExchangeUsecase(a DCEInterfaces.IDailyCurrencyExchangeRepository,b CEInterfaces.ICurrencyExchangeRepository) DCEInterfaces.IDailyCurrencyExchangeUseCase {
	return &DCEUsecase{
		DCERepo: a,
		CERepo: b,
	}
}

func (a *DCEUsecase) Store(c *gin.Context,dce *models.DailyCurrencyExchange) (*models.DailyCurrencyExchange, error) {

	exist,dailyCur, _ := a.DCERepo.FindAvailable(dce)

	if exist {
		//update
		res, err := a.DCERepo.Update(dce,dailyCur)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	res, err := a.DCERepo.Store(dce)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *DCEUsecase) GetCurrencyExchangeID(c *gin.Context,m *models.CurrencyExchange) (int, error) {
	id, err := a.CERepo.FindByCurrency(m)
	if err != nil {
		return 0, err
	}

	if id <= 0 {
		return 0,errors.New("Data Not Found")
	}

	return id, nil
}

func (a *DCEUsecase) FetchByDate(c *gin.Context,date string) ([]*models.ListDailyCurrency, *models.Pagination, error) {
	
	if date == "" {
		return nil,nil,errors.New("parameter date not be null")
	}

	var (
		pagination *models.Pagination
		total      int
	)
	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("perpage"))

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 10
	}

	//Get All
	currencyList, total, err := a.DCERepo.FetchByDate(page, perPage, date)

	if err != nil {
		return nil, nil, err
	}

	count := len(currencyList)
	pagination = models.BuildPagination(total, page, perPage, count)

	return currencyList, pagination, err
}

func (a *DCEUsecase) FetchTrendCurrency(c *gin.Context,from string,to string) ([]*models.ListTrendDailyCurrency,float64, error) {

	var variance float64
	
	if from == "" || to == "" {
		return nil,0,errors.New("parameter from or to not be null")
	}

	currencyExchange := mapCurrencyData(c);
	id, err := a.CERepo.FindByCurrency(currencyExchange)

	//Get All
	currencyTrendList, err := a.DCERepo.FetchTrendCurrency(id)

	if err != nil {
		return nil,0, err
	}

	countTrend := len(currencyTrendList)

	if countTrend <= 0 {
		return nil,0, errors.New("insufficient data")
	}
	variance = currencyTrendList[0].MaxRate
	if countTrend > 1 {
		for _, curTrend := range currencyTrendList {
			maxRate := curTrend.MaxRate
			minRate := curTrend.MinRate

			variance = maxRate - minRate
			break;
		}
	}

	return currencyTrendList,variance, err
}

func mapCurrencyData(params *gin.Context) *models.CurrencyExchange{
	curExchange := new(models.CurrencyExchange)
	curExchange.From = params.Query("from")
	curExchange.To 	 = params.Query("to")
	return curExchange
}