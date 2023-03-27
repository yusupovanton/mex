package dto

import "time"

type BinanceResponse struct {
	Code          string      `json:"code"`
	Message       interface{} `json:"message"`
	MessageDetail interface{} `json:"messageDetail"`
	Data          []*Adv      `json:"data"`
	Total         int         `json:"total"`
	Success       bool        `json:"success"`
}

type Adv struct {
	AdInfo Advertisement `json:"adv"`
}

type Advertiser struct {
	UserNo           string      `json:"userNo"`
	RealName         interface{} `json:"realName"`
	NickName         string      `json:"nickName"`
	Margin           interface{} `json:"margin"`
	MarginUnit       interface{} `json:"marginUnit"`
	OrderCount       interface{} `json:"orderCount"`
	MonthOrderCount  int         `json:"monthOrderCount"`
	MonthFinishRate  float64     `json:"monthFinishRate"`
	AdvConfirmTime   interface{} `json:"advConfirmTime"`
	Email            interface{} `json:"email"`
	RegistrationTime interface{} `json:"registrationTime"`
	Mobile           interface{} `json:"mobile"`
	UserType         string      `json:"userType"`
	TagIconUrls      []string    `json:"tagIconUrls"`
	UserGrade        int         `json:"userGrade"`
	UserIdentity     string      `json:"userIdentity"`
	ProMerchant      interface{} `json:"proMerchant"`
	IsBlocked        interface{} `json:"isBlocked"`
}

type Advertisement struct {
	AdvNo              string      `json:"advNo"`
	Classify           string      `json:"classify"`
	TradeType          string      `json:"tradeType"`
	Asset              string      `json:"asset" db:"asset"`
	FiatUnit           string      `json:"fiatUnit"`
	AdvStatus          interface{} `json:"advStatus"`
	PriceType          interface{} `json:"priceType"`
	PriceFloatingRatio interface{} `json:"priceFloatingRatio"`
	RateFloatingRatio  interface{} `json:"rateFloatingRatio"`
	CurrencyRate       interface{} `json:"currencyRate"`
	Price              string      `json:"price" db:"price"`
	InitAmount         interface{} `json:"initAmount"`
	SurplusAmount      string      `json:"surplusAmount"`
}

type BinanceDBRow struct {
	AdvNo     string    `db:"adv_no"`
	Asset     string    `db:"asset"`
	Price     string    `db:"price"`
	UpdatedAt time.Time `db:"updated_at"`
}
