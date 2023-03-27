package service

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yusupovanton/moneyExchange/internal/me-scraper/app/dto"
)

type ScraperService struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewScraperService(db *sqlx.DB, ctx context.Context) *ScraperService {
	return &ScraperService{
		db:  db,
		ctx: ctx,
	}
}

func (s *ScraperService) ScrapeBinance() error {

	now := time.Now()
	url := "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"
	method := "POST"
	payload := strings.NewReader(`{"proMerchantAds":false,"page":1,"rows":10,"payTypes":["RaiffeisenBank"],"countries":[],"publisherType":null,"asset":"USDT","fiat":"RUB","tradeType":"BUY"}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Printf("Could not create request: %v", err)
		return err
	}

	req.Header.Add("Referer", "https://p2p.binance.com/en/trade/RaiffeisenBank/USDT?fiat=RUB")
	req.Header.Add("Cookie", "cid=ut3e6vw9; c2c-menu-ssbt=false; c2c-menu-ssct=false; _ga=GA1.2.1720100647.1679855382; _gid=GA1.2.1146274603.1679855382; common_fiat=RUB; OptanonAlertBoxClosed=2023-03-26T18:29:54.781Z; OptanonConsent=isGpcEnabled=0&datestamp=Sun+Mar+26+2023+23%3A43%3A26+GMT%2B0300+(GMT%2B03%3A00)&version=202211.1.0&isIABGlobal=false&hosts=&consentId=4e4005c7-7929-407d-898f-877e48cce728&interactionCount=1&groups=C0001%3A1%2CC0003%3A0%2CC0004%3A0%2CC0002%3A0&landingPath=NotLandingPage&geolocation=TR%3B35&AwaitingReconsent=false; _cq_duid=1.1679855384.FOkkiQvF73EPHLyQ; _cq_suid=1.1679860661.NlWq8MRQipS5AM2U; _gat_UA-162512367-1=1; campaign=www.google.com; source=organic; userPreferredCurrency=USD_USD; sys_mob=no; fiat-prefer-currency=EUR; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%221871f2f009099a-01f0329161f9add-3c626b4b-1296000-1871f2f00911048%22%2C%22first_id%22%3A%22%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E8%87%AA%E7%84%B6%E6%90%9C%E7%B4%A2%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC%22%2C%22%24latest_referrer%22%3A%22https%3A%2F%2Fwww.google.com%2F%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMTg3MWYyZjAwOTA5OWEtMDFmMDMyOTE2MWY5YWRkLTNjNjI2YjRiLTEyOTYwMDAtMTg3MWYyZjAwOTExMDQ4In0%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%22%2C%22value%22%3A%22%22%7D%2C%22%24device_id%22%3A%221871f2f009099a-01f0329161f9add-3c626b4b-1296000-1871f2f00911048%22%7D; _ga_3WP50LGEEC=GS1.1.1679860718.1.0.1679860720.58.0.0; _uetsid=997e9c10cc1011edae1b810c76ea286d; _uetvid=997ea490cc1011ed8ee98f1985b689b1; lang=en; se_gd=wcGFBA18THHEAkH4HDQ4gZZAxVFtQBVUFAXdfU09VhcVQFVNXVYV1; se_gsd=BionLBl/LCokUFo3JCYiFQAiDQoIAwcGUFVEVVRVVFhaNFNS1; se_sd=RQIFBVF8SQXUxMXIXBFAgZZHRFQELESW1sQdfU09VhcVQEFNXV0Q1; sajssdk_2015_cross_new_user=1; BNC_FV_KEY=33a632abe210ad9e93458f77bfa4d80adba8682d; BNC_FV_KEY_EXPIRE=1679876985698; _gcl_au=1.1.2133346388.1679855384; bnc-uuid=6fa5ceb4-df87-46dd-a746-7be6a95bda5f")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Safari/605.1.15")
	req.Header.Add("Host", "p2p.binance.com")
	req.Header.Add("Origin", "https://p2p.binance.com")
	req.Header.Add("Content-Length", "155")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.9")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("lang", "en")
	req.Header.Add("clienttype", "web")
	req.Header.Add("bnc-uuid", "6fa5ceb4-df87-46dd-a746-7be6a95bda5f")
	req.Header.Add("x-trace-id", "e23183d6-a968-4343-96b3-2de4c0fcf41a")
	req.Header.Add("x-ui-request-trace", "e23183d6-a968-4343-96b3-2de4c0fcf41a")
	req.Header.Add("fvideo-id", "33a632abe210ad9e93458f77bfa4d80adba8682d")
	req.Header.Add("device-info", "eyJzY3JlZW5fcmVzb2x1dGlvbiI6IjkwMCwxNDQwIiwiYXZhaWxhYmxlX3NjcmVlbl9yZXNvbHV0aW9uIjoiODAxLDE0NDAiLCJzeXN0ZW1fdmVyc2lvbiI6Ik1hYyBPUyAxMC4xNS43IiwiYnJhbmRfbW9kZWwiOiJ1bmtub3duIiwic3lzdGVtX2xhbmciOiJlbi1HQiIsInRpbWV6b25lIjoiR01UKzMiLCJ0aW1lem9uZU9mZnNldCI6LTE4MCwidXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0LzYwNS4xLjE1IChLSFRNTCwgbGlrZSBHZWNrbykgVmVyc2lvbi8xNi4yIFNhZmFyaS82MDUuMS4xNSIsImxpc3RfcGx1Z2luIjoiV2ViS2l0IGJ1aWx0LWluIFBERiIsImNhbnZhc19jb2RlIjoiMzg4MWRlNmIiLCJ3ZWJnbF92ZW5kb3IiOiJBcHBsZSBJbmMuIiwid2ViZ2xfcmVuZGVyZXIiOiJBcHBsZSBHUFUiLCJhdWRpbyI6IjEyNC4wNDM0NTgwODg3Mzc2OCIsInBsYXRmb3JtIjoiTWFjSW50ZWwiLCJ3ZWJfdGltZXpvbmUiOiJFdXJvcGUvSXN0YW5idWwiLCJkZXZpY2VfbmFtZSI6IlNhZmFyaSBWMTYuMiAoTWFjIE9TKSIsImZpbmdlcnByaW50IjoiMjQ4ZWU0OTI2YWMyMzRhZDRkMDJhODkwZDBkNTgxMTciLCJkZXZpY2VfaWQiOiIiLCJyZWxhdGVkX2RldmljZV9pZHMiOiIifQ==")
	req.Header.Add("c2ctype", "c2c_merchant")
	req.Header.Add("csrftoken", "d41d8cd98f00b204e9800998ecf8427e")

	res, err := client.Do(req)

	if err != nil {
		log.Printf("Could not send the request: %v", err)
		return err
	}

	defer res.Body.Close()

	reader, err := gzip.NewReader(res.Body)
	if err != nil {
		log.Printf("Could not decompress body of response: %v", err)
		return err
	}
	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		log.Printf("Could not read body of response: %v", err)
		return err
	}

	resp := dto.BinanceResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Printf("Could not unmarshall body of response to structs: %v", err)
		return err
	}

	rowsToInsert := responseToRows(resp.Data, now)

	err = s.insertToDb(s.ctx, rowsToInsert)

	if err != nil {
		log.Printf("Could not post data to db: %v", err)
		return err
	}

	log.Println("Finished! Going to sleep for 10s")
	time.Sleep(time.Second * 10)

	return nil
}

func responseToRows(advArr []*dto.Adv, now time.Time) []*dto.BinanceDBRow {

	var convertedRows = []*dto.BinanceDBRow{}

	for _, adv := range advArr {
		row := dto.BinanceDBRow{}

		row.AdvNo = adv.AdInfo.AdvNo
		row.Asset = adv.AdInfo.Asset
		row.Price = adv.AdInfo.Price
		row.UpdatedAt = now

		convertedRows = append(convertedRows, &row)
	}

	return convertedRows
}

var queryInsert = `
INSERT INTO public.me_active_ads(
	adv_no,
	asset, 
	price,
	updated_at
) VALUES (
	:adv_no, 
	:asset,
	:price,
	:updated_at
) ON CONFLICT 
	(adv_no)
DO UPDATE SET
	asset = EXCLUDED.asset,
	price = EXCLUDED.price,
	updated_at = EXCLUDED.updated_at
`

func (s *ScraperService) insertToDb(ctx context.Context, rowsarr []*dto.BinanceDBRow) error {

	_, err := s.db.NamedExecContext(ctx, queryInsert, rowsarr)

	if err != nil {
		log.Printf("An error occured while executing query: %v", err)
		return err
	}

	return nil
}

// Tinkoff add + field add to the db
