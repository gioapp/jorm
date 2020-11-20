package main

import (
	"fmt"
	"gioui.org/app"
	l "gioui.org/layout"
	jorm "github.com/gioapp/jorm/app"
	"github.com/gioapp/jorm/cfg"
	in "github.com/gioapp/jorm/cfg/ini"
	"github.com/gioapp/jorm/mod/c"
	csrc "github.com/gioapp/jorm/mod/c/src"
	"github.com/p9c/pod/pkg/gui/f"
	"os"
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
	quit := make(chan struct{})

	j := jorm.NewJorm()

	j.Coins = c.ReadAllCoins()

	go func() {
		if err := f.NewWindow().
			Size(800, 600).
			Title("app").
			Open().
			Run(j.Jorm,
				func(gtx l.Context) {},
				func() {
					close(quit)
					os.Exit(0)
				}, quit); Check(err) {
		}
	}()
	app.Main()

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
	//}()

	// exp.SrcNode().GetAddrs()

}
