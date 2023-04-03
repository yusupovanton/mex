package dto

import "time"

type CurrencyExchangeRate struct {
	Success   bool               `json:"success"`
	Timestamp int64              `json:"timestamp" db:"ts"`
	Base      string             `json:"base" db: "base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

type CurrencyDBRow struct {
	Ts          time.Time `db:"ts"`
	Base        string    `db:"base"`
	Euro        float64   `db:"euro"`
	UsDollar    float64   `db:"us_dollar"`
	TurkishLira float64   `db:"turkish_lira"`
}
