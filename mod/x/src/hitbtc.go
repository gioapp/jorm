package xsrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gioapp/jorm/jdb"
	"github.com/gioapp/jorm/mod/x"
)

type HitBTCExchange struct {
	ID                   string `json:"id"`
	BaseCurrency         string `json:"baseCurrency"`
	QuoteCurrency        string `json:"quoteCurrency"`
	QuantityIncrement    string `json:"quantityIncrement"`
	TickSize             string `json:"tickSize"`
	TakeLiquidityRate    string `json:"takeLiquidityRate"`
	ProvideLiquidityRate string `json:"provideLiquidityRate"`
	FeeCurrency          string `json:"feeCurrency"`
}

type HitBTCExchangeTickers struct {
	Ask         string    `json:"ask"`
	Bid         string    `json:"bid"`
	Last        string    `json:"last"`
	Open        string    `json:"open"`
	Low         string    `json:"low"`
	High        string    `json:"high"`
	Volume      string    `json:"volume"`
	VolumeQuote string    `json:"volumeQuote"`
	Timestamp   time.Time `json:"timestamp"`
	Symbol      string    `json:"symbol"`
}

func getHitBTCExchange() {
	fmt.Println("GetHitBTCExchangeStart")
	exchangeRaw := []HitBTCExchange{}
	tickersRaw := []HitBTCExchangeTickers{}

	slug := "hitbtc"
	var exchange x.Exchange
	exchange.Name = "HitBTC"
	exchange.Slug = slug
	resps, err := http.Get("https://api.hitbtc.com/api/2/public/symbol")
	if err != nil {
	}
	defer resps.Body.Close()
	mapBodyS, err := ioutil.ReadAll(resps.Body)
	json.Unmarshal(mapBodyS, &exchangeRaw)

	respt, err := http.Get("https://api.hitbtc.com/api/2/public/ticker")
	if err != nil {
	}
	defer respt.Body.Close()
	mapBodyT, err := ioutil.ReadAll(respt.Body)
	json.Unmarshal(mapBodyT, &tickersRaw)
	tickers := make(map[string]HitBTCExchangeTickers)

	for _, ticker := range tickersRaw {
		tickers[ticker.Symbol] = ticker
	}

	markets := make(map[string][]x.Currency)
	for _, marketSrc := range exchangeRaw {
		cur := x.Currency{
			Symbol: marketSrc.BaseCurrency,
			Ask:    tickers[marketSrc.ID].Ask,
			Bid:    tickers[marketSrc.ID].Bid,
			High:   tickers[marketSrc.ID].High,
			Last:   tickers[marketSrc.ID].Last,
			Low:    tickers[marketSrc.ID].Low,
			Volume: tickers[marketSrc.ID].Volume,
		}
		_, ok := markets[marketSrc.QuoteCurrency]
		if !ok {
			markets[marketSrc.QuoteCurrency] = []x.Currency{}
		}
		markets[marketSrc.QuoteCurrency] = append(markets[marketSrc.QuoteCurrency], cur)
	}
	var marketsSlice x.Markets
	for i := range markets {
		newSlice := x.Market{
			Symbol:     i,
			Currencies: markets[i],
		}
		marketsSlice = append(marketsSlice, newSlice)
	}
	exchange.Markets = marketsSlice
	jdb.WriteExchange(slug, exchange)

	fmt.Println("GetHitBTCExchangeDone")

}
