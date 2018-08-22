package usecase

import (
	"errors"
	CEInterfaces "exchange-currency/app/currency-exchange"
	"exchange-currency/models"
	"strconv"
	"github.com/gin-gonic/gin"
)

type CEUsecase struct {
	CERepo CEInterfaces.ICurrencyExchangeRepository
}

func NewCurrencyExchangeUsecase(a CEInterfaces.ICurrencyExchangeRepository) CEInterfaces.ICurrencyExchangeUseCase {
	return &CEUsecase{
		CERepo: a,
	}
}

func (a *CEUsecase) Store(c *gin.Context, ce *models.CurrencyExchange) (*models.CurrencyExchange, error) {

	exist, err := a.CERepo.FindAvailable(ce)

	if err != nil {
		return nil, err
	}

	if exist {
		return nil, errors.New("Currency exchange already exist")
	}

	ce.Status = models.STATUS_ENABLE
	res, err := a.CERepo.Store(ce)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *CEUsecase) Fetch(c *gin.Context) ([]*models.CurrencyExchange, *models.Pagination, error) {
	var (
		pagination *models.Pagination
		total      int
	)

	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("perpage"))
	status := c.Query("status")

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 10
	}

	//Get All
	currencies, total, err := a.CERepo.Fetch(page, perPage, status)

	if err != nil {
		return nil, nil, err
	}

	count := len(currencies)
	pagination = models.BuildPagination(total, page, perPage, count)

	return currencies, pagination, err
}

func (a *CEUsecase) Delete(c *gin.Context, id uint64) error {

	if id == 0 {
		return errors.New("parameter id not be null")
	}

	modifiedBy, _ := strconv.Atoi(c.Param("modified_by"))

	m, err := a.CERepo.GetById(id)

	if err != nil {
		return err
	}

	errDel := a.CERepo.Delete(m, modifiedBy)
	if errDel != nil {
		return errDel
	}
	return nil
}
