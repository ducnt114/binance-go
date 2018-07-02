package binance_go

const (
	BTCSymbol = "BTC"
	ETHSymbol = "ETH"
)

type SymbolInfo struct {
	Symbol     string `json:"symbol"`
	Status     string `json:"trading"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
}

type ExchangeInfo struct {
	TimeZone   string        `json:"timezone"`
	ServerTime int64         `json:"serverTime"`
	Symbols    []*SymbolInfo `json:"symbols"`
}

type CandlestickData struct {
	OpenTime                 float64
	OpenPrice                float64
	HighPrice                float64
	LowPrice                 float64
	ClosePrice               float64
	Volume                   float64
	CloseTime                float64
	QuoteAssetVolume         float64
	NumberOfTrades           float64
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
	Ignore                   float64
}
