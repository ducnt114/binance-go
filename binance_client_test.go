package binance_go

import (
	"testing"
	"fmt"
	"time"
)

func TestBinanceClient_GetListSymbol(t *testing.T) {
	client := NewBinanceClient()

	symbols, err := client.GetListSymbol()
	if err != nil {
		t.Fatal(err)
	}

	counter := 0
	for _, s := range symbols {
		if s.QuoteAsset == BTCSymbol {
			counter += 1
			fmt.Println(s.Symbol)
		}
	}
	fmt.Println("Size: ", counter)
}

func TestBinanceClient_GetCandlestickData(t *testing.T) {
	client := NewBinanceClient()

	startTime := 1000 * (time.Now().Unix() - 3*24*60*60)
	endTime := time.Now().Unix() * 1000

	data, err := client.GetCandlestickData("ETHUSDT", Interval1d, startTime, endTime)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Data size: ", len(data))

	for _, d := range data {
		fmt.Println("***")
		fmt.Println(d.OpenTime)
		fmt.Println(d.OpenPrice)
		fmt.Println(d.HighPrice)
		fmt.Println(d.LowPrice)
		fmt.Println(d.ClosePrice)
		fmt.Println(d.Volume)
		fmt.Println(d.CloseTime)
		fmt.Println(d.QuoteAssetVolume)
		fmt.Println(d.NumberOfTrades)
		fmt.Println(d.TakerBuyBaseAssetVolume)
		fmt.Println(d.TakerBuyQuoteAssetVolume)
		fmt.Println(d.Ignore)
		fmt.Println("***")
	}
}
