package binance_go

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

const (
	baseEndpoint = "https://api.binance.com"

	Interval1m  = "1m"
	Interval3m  = "3m"
	Interval5m  = "5m"
	Interval15m = "15m"
	Interval30m = "30m"
	Interval1h  = "1h"
	Interval2h  = "2h"
	Interval4h  = "4h"
	Interval6h  = "6h"
	Interval8h  = "8h"
	Interval12h = "12h"
	Interval1d  = "1d"
	Interval3d  = "3d"
	Interval1w  = "1w"
	Interval1M  = "1M"
)

type binanceClient struct {
	BaseEndpoint string
}

func NewBinanceClient() *binanceClient {
	return &binanceClient{
		BaseEndpoint: baseEndpoint,
	}
}

func (b *binanceClient) GetListSymbol() ([]*SymbolInfo, error) {
	res := make([]*SymbolInfo, 0)
	getURL := fmt.Sprintf("%s/api/v1/exchangeInfo", b.BaseEndpoint)

	resp, err := http.Get(getURL)
	if err != nil {
		log.Println(err)
		return res, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return res, err
	}

	exchangeInfo := &ExchangeInfo{}
	err = json.Unmarshal(respBytes, &exchangeInfo)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return exchangeInfo.Symbols, nil
}

func (b *binanceClient) GetCandlestickData(symbol, interval string, startTime, endTime int64) (
	[]*CandlestickData, error) {

	res := make([]*CandlestickData, 0)

	getURL := fmt.Sprintf("%s/%s?symbol=%s&interval=%s&startTime=%d&endTime=%d",
		b.BaseEndpoint,
		"api/v1/klines",
		symbol,
		interval,
		startTime,
		endTime)

	resp, err := http.Get(getURL)
	if err != nil {
		log.Println(err)
		return res, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return res, err
	}

	fmt.Println("Body: ", string(respBytes))

	var listKlines []interface{} = make([]interface{}, 0)
	err = json.Unmarshal(respBytes, &listKlines)
	if err != nil {
		log.Println(err)
		return res, err
	}

	for _, kline := range listKlines {
		klines := kline.([]interface{})

		candlestickData := &CandlestickData{}

		candlestickData.OpenTime, _ = klines[0].(float64)
		candlestickData.OpenPrice, _ = strconv.ParseFloat(klines[1].(string), 64)
		candlestickData.HighPrice, _ = strconv.ParseFloat(klines[2].(string), 64)
		candlestickData.LowPrice, _ = strconv.ParseFloat(klines[3].(string), 64)
		candlestickData.ClosePrice, _ = strconv.ParseFloat(klines[4].(string), 64)
		candlestickData.Volume, _ = strconv.ParseFloat(klines[5].(string), 64)
		candlestickData.CloseTime, _ = klines[6].(float64)
		candlestickData.QuoteAssetVolume, _ = strconv.ParseFloat(klines[7].(string), 64)
		candlestickData.NumberOfTrades, _ = klines[8].(float64)
		candlestickData.TakerBuyBaseAssetVolume, _ = strconv.ParseFloat(klines[9].(string), 64)
		candlestickData.TakerBuyQuoteAssetVolume, _ = strconv.ParseFloat(klines[10].(string), 64)
		candlestickData.Ignore, _ = strconv.ParseFloat(klines[11].(string), 64)

		res = append(res, candlestickData)
	}

	return res, err
}
