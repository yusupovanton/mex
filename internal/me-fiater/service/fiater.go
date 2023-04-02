package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/yusupovanton/moneyExchange/internal/me-fiater/app/dto"
)

type FiaterService struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewFiaterService(db *sqlx.DB, ctx context.Context) *FiaterService {

	return &FiaterService{
		db:  db,
		ctx: ctx,
	}
}

func (s *FiaterService) GetFiatConversionRates() {
	url := "https://api.apilayer.com/fixer/latest?base=RUB&symbols=USD,TRY"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("apikey", "x7qU0fh3Vhr24wDoB0pGhjT8XceaDGaL")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &dto.CurrencyExchangeRate{})

	if err != nil {
		fmt.Println(err)
		return
	}

}
