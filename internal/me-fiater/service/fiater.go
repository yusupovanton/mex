package service

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

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

func (s *FiaterService) GetFiatConversionRates() error {
	url := "https://api.apilayer.com/exchangerates_data/latest?base=RUB&symbols=USD,EUR,TRY"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("apikey", "x7qU0fh3Vhr24wDoB0pGhjT8XceaDGaL")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	jsonRes := &dto.CurrencyExchangeRate{}

	err = json.Unmarshal(body, jsonRes)

	if err != nil {
		return err
	}
	row := s.resultToRow(jsonRes)
	s.postResultToDb(s.ctx, row)
	return err
}

func (s *FiaterService) resultToRow(resJson *dto.CurrencyExchangeRate) *dto.CurrencyDBRow {

	euro := resJson.Rates["EUR"]
	dollar := resJson.Rates["USD"]
	lira := resJson.Rates["TRY"]

	return &dto.CurrencyDBRow{
		Ts:          time.Unix(resJson.Timestamp, 0),
		Base:        resJson.Base,
		Euro:        euro,
		UsDollar:    dollar,
		TurkishLira: lira,
	}
}

const queryInsert = `
INSERT INTO public.me_fiat_conversion_rates(
	base,
	us_dollar,
	euro,
	turkish_lira,
	ts
) VALUES (
	:base,
	:us_dollar,
	:euro,
	:turkish_lira,
	:ts
);
`

func (s *FiaterService) postResultToDb(ctx context.Context, row *dto.CurrencyDBRow) error {

	_, err := s.db.NamedExecContext(ctx, queryInsert, row)

	if err != nil {
		log.Printf("An error occured while executing query: %v", err)
		return err
	}

	return nil
}
