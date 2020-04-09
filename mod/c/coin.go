package c

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"strings"

	"github.com/gioapp/jorm/cfg"
	"github.com/gioapp/jorm/jdb"
	"github.com/gioapp/jorm/mod"
)

type Coins struct {
	N int    `json:"n"`
	C []Coin `json:"c"`
}
type CoinsBase struct {
	N int        `json:"n"`
	C []CoinBase `json:"c"`
}

// Coin stores identifying information about coins in the database
type CoinBase struct {
	Rank int    `json:"r"`
	Name string `json:"n"`
	Slug string `json:"s"`
}

type Coin struct {
	Name   string      `json:"n" form:"name"`
	Ticker string      `json:"t" form:"ticker"`
	Slug   string      `json:"s" form:"slug"`
	Logo   image.Image `json:"l" form:"logo"`
}

// CoinData stores all of the information relating to a coin
type CoinData struct {
	Rank    int    `json:"rank" form:"rank"`
	Name    string `json:"name" form:"name"`
	Ticker  string `json:"symbol" form:"symbol"`
	Slug    string `json:"slug" form:"slug"`
	Algo    string `json:"algo" form:"algo"`
	BitNode bool   `json:"bitnode" form:"bitnode"`

	Token bool `json:"token" form:"token"`
	Ico   bool `json:"ico" form:"ico"`

	Description          string `json:"desc"`
	WebSite              string `json:"web"`
	TotalCoinSupply      string `json:"total"`
	DifficultyAdjustment string `json:"diff"`
	BlockRewardReduction string `json:"rew"`
	ProofType            string `json:"proof"`
	StartDate            string `json:"start"`

	Twitter string `json:"tw"`
	// Explorers            []string `json:"explorers"`
	// Boards               []string `json:"boards"`
	Facebook   string `json:"facebook"`
	Reddit     string `json:"reddit"`
	Github     string `json:"github"`
	WhitePaper string `json:"whitepaper"`
	Published  bool   `json:"published" form:"published"`
}

type ICO struct {
	Status                      string `json:"Status"`
	Description                 string `json:"Description"`
	TokenType                   string `json:"TokenType"`
	Website                     string `json:"Website"`
	PublicPortfolioURL          string `json:"PublicPortfolioUrl"`
	PublicPortfolioID           string `json:"PublicPortfolioId"`
	Features                    string `json:"Features"`
	FundingTarget               string `json:"FundingTarget"`
	FundingCap                  string `json:"FundingCap"`
	ICOTokenSupply              string `json:"ICOTokenSupply"`
	TokenSupplyPostICO          string `json:"TokenSupplyPostICO"`
	TokenPercentageForInvestors string `json:"TokenPercentageForInvestors"`
	TokenReserveSplit           string `json:"TokenReserveSplit"`
	Date                        int    `json:"Date"`
	EndDate                     int    `json:"EndDate"`
	FundsRaisedList             string `json:"FundsRaisedList"`
	FundsRaisedUSD              string `json:"FundsRaisedUSD"`
	StartPrice                  string `json:"StartPrice"`
	StartPriceCurrency          string `json:"StartPriceCurrency"`
	PaymentMethod               string `json:"PaymentMethod"`
	Jurisdiction                string `json:"Jurisdiction"`
	LegalAdvisers               string `json:"LegalAdvisers"`
	LegalForm                   string `json:"LegalForm"`
	SecurityAuditCompany        string `json:"SecurityAuditCompany"`
	Blog                        string `json:"Blog"`
	WhitePaper                  string `json:"WhitePaper"`
	WhitePaperLink              string `json:"WhitePaperLink"`
}

// ReadAllCoins reads in all of the data about all coins in the database
func ReadAllCoins() Coins {
	coins := jdb.ReadData(cfg.Web + "/coins")
	cs := make([]Coin, len(coins))
	csb := CoinsBase{}

	csb.N = 0
	for i := range coins {
		csb.N++
		if err := json.Unmarshal(coins[i], &cs[i]); err != nil {
			fmt.Println("Error", err)
		}

		// Load logo image from database
		l := jdb.Read("data/"+cs[i].Slug, "logo")
		logos := l.(map[string]interface{})

		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(logos["img16"].(string)))
		logo, _, err := image.Decode(reader)
		if err != nil {
			//log.Fatal(err)
		}

		ccb := CoinBase{
			Rank: csb.N,
			Name: cs[i].Name,
			Slug: cs[i].Slug,
		}
		cs[i].Logo = logo
		csb.C = append(csb.C, ccb)
	}

	cns := Coins{
		N: csb.N,
		C: cs,
	}
	c := mod.Cache{Data: cns}
	cb := mod.Cache{Data: csb}
	jdb.DB.Write(cfg.Web, "coins", c)
	jdb.DB.Write(cfg.Web, "coinsbase", cb)
	return cns
}
