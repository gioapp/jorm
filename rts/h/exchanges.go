package h

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.parallelcoin.io/dev/jorm/cfg"
	"git.parallelcoin.io/dev/jorm/jdb"
	"git.parallelcoin.io/dev/jorm/mod/c"
	"git.parallelcoin.io/dev/jorm/mod/x"
	"github.com/gorilla/mux"
)

func ViewMarket(w http.ResponseWriter, r *http.Request) {
	rc := mux.Vars(r)["coin"]
	var coin c.Coin
	var coinMarkets x.CoinMarkets
	jdb.DB.Read(cfg.Web+"/coins", rc, &coin)
	exchanges := x.ReadAllExchanges()
	for _, exchange := range exchanges {
		for _, market := range exchange.Markets {
			for _, cur := range market.Currencies {
				if cur.Symbol == coin.Ticker {
					coinMarket := x.CoinMarket{
						Exchange:     exchange.Name,
						ExchangeSlug: exchange.Slug,
						Market:       market.Symbol,
						Ticker:       cur,
					}
					coinMarkets = append(coinMarkets, coinMarket)
				}
			}
		}
	}
	x := map[string]interface{}{
		"d": coinMarkets,
	}
	out, err := json.Marshal(x)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
