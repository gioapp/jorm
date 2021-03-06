package h

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"git.parallelcoin.io/dev/utl"
	"github.com/gioapp/jorm/cfg"
	"github.com/gioapp/jorm/jdb"
	"github.com/gioapp/jorm/mod"
	"github.com/gioapp/jorm/mod/a"
	"github.com/gioapp/jorm/mod/c"
	"github.com/gioapp/jorm/tpl"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/json"
)

type home struct {
	D []c.Coin
	C c.Coins
}

// HomeHandler handles a request for (?)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	coins := c.ReadAllCoins()
	var bitnoded []c.Coin
	for _, coin := range coins.C {

		if utl.FileExists(cfg.Web + "/data/" + coin.Slug + "/bitnodes.json") {
			bitnoded = append(bitnoded, coin)
		}
	}
	data := home{
		D: bitnoded,
		C: coins,
	}
	tpl.TemplateHandler().ExecuteTemplate(w, "base_gohtml", data)
}

// AddCoinHandler handles a request for adding coin data
func AddCoinHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("coin")
	coin := c.Coin{
		Name: name,
		Slug: utl.MakeSlug(name),
	}

	//fmt.Println("name", name)
	//fmt.Println("coin", coin)
	c := mod.Cache{Data: coin}
	jdb.DB.Write(cfg.Web+"/coins", coin.Slug, c)
	http.Redirect(w, r, "/", 302)
}

// AddNodeHandler handles a request for adding node data
func AddNodeHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.FormValue("ip")
	p := r.FormValue("port")
	c := utl.MakeSlug(r.FormValue("coin"))

	var bitNodes a.BitNodes
	jdb.DB.Read(c, "bitnodes", &bitNodes)

	port, err := strconv.ParseInt(p, 10, 64)
	if err == nil {
		// What is this supposed to be printing exactly?
		// fmt.Printf("%d of type %T", p, p)
	}

	fmt.Println("ip", ip)
	fmt.Println("port", port)

	bitNode := a.BitNode{
		IP:   ip,
		Port: port,
	}

	bitNodes = append(bitNodes, bitNode)
	b := mod.Cache{Data: bitNodes}
	jdb.DB.Write(c, "bitnodes", b)
	http.Redirect(w, r, "/", 302)
}

// CoinsHandler handles a request for coin data
func CoinsHandler(w http.ResponseWriter, r *http.Request) {
	// c := map[string]interface{}{
	// 	"d": map[string]interface{}{
	// 		// "coins": Coins(),
	// 	},
	// }
	// out, err := json.Marshal(c)
	// if err != nil {
	// 	fmt.Println("Error encoding JSON")
	// 	return
	// }
	// w.Write([]byte(out))
}

// CoinNodesHandler handles a request for (?)
func CoinNodesHandler(w http.ResponseWriter, r *http.Request) {

}

// NodeHandler handles a request for (?)
func NodeHandler(w http.ResponseWriter, r *http.Request) {

}

// NodeHandler handles a request for (?)
func ViewJSON() http.Handler {
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)

	return http.StripPrefix("/j", m.Middleware(http.FileServer(http.Dir(cfg.Dir+cfg.Web))))
}
