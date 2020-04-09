package main

import (
	"fmt"
	"gioui.org/app"
	"github.com/gioapp/jorm/cfg"
	"github.com/gioapp/jorm/cfg/ini"
	"github.com/gioapp/jorm/jorm"
	"github.com/gioapp/jorm/mod/c"
	csrc "github.com/gioapp/jorm/mod/c/src"
	"time"
)

func main() {
	//coins := c.Coins{}
	in.Init()
	if cfg.Initial {
		fmt.Println("running initial sync")
		csrc.GetCoinSources()
		time.Sleep(time.Second * 2)
	}

	go func() {
		j := jorm.NewJorm()
		j.Coins = c.ReadAllCoins()
		//r := rts.Routes()

		//go func() {
		//}()
		//cr := cron.New()
		//cr.AddFunc("@every 60s", func() {
		//	fmt.Println("Radi kron GetBitNodes")
		//	n.GetBitNodes(coins)
		//})
		//
		//cr.AddFunc("@every 5555s", func() {
		//	csrc.GetCoinSources()
		//})
		//cr.AddFunc("@every 60s", func() {
		//	xsrc.GetExchangeSources()
		//	// dsrc.GetDataSources()
		//})
		//cr.Start()

		//go func() {
		//	go log.Fatal(http.ListenAndServe(":3111", r))
		//}()
		j.Gui()
	}()
	app.Main()

	// exp.SrcNode().GetAddrs()

}
