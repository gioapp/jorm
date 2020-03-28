package rts

import (
	"git.parallelcoin.io/dev/jorm/rts/h"
	"github.com/gorilla/mux"
)

// SetUpRoutes sets up routes for jorm
func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", h.HomeHandler)

	f := r.PathPrefix("/f").Subrouter()
	f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
	f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

	// a := r.PathPrefix("/a").Subrouter()
	// a.HandleFunc("/coins", CoinsHandler).Methods("GET")
	// a.HandleFunc("/{coin}/nodes", CoinNodesHandler).Methods("GET")
	// a.HandleFunc("/{coin}/{nodeip}", NodeHandler).Methods("GET")

	// s := r.PathPrefix("/s").Subrouter()
	// s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./tpl/static/"))))

	b := r.PathPrefix("/b").Subrouter()
	b.HandleFunc("/{coin}/blocks/{per}/{page}", h.ViewBlocks).Methods("GET")
	b.HandleFunc("/{coin}/lastblock", h.LastBlock).Methods("GET")
	// b.HandleFunc("/{coin}/block/{blockheight}", h.ViewBlockHeight).Methods("GET")
	b.HandleFunc("/{coin}/block/{blockheight}", h.ViewBlockHeight).Methods("GET")
	b.HandleFunc("/{coin}/hash/{blockhash}", h.ViewHash).Methods("GET")
	b.HandleFunc("/{coin}/tx/{txid}", h.ViewTx).Methods("GET")

	b.HandleFunc("/{coin}/market", h.ViewMarket).Methods("GET")

	j := r.PathPrefix("/j").Subrouter()
	// j.HandleFunc("/", h.ViewJSON).Methods("GET")
	j.PathPrefix("/").Handler(h.ViewJSON())

	return r
}
