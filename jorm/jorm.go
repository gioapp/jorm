package jorm

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/unit"
	"git.parallelcoin.io/dev/jorm/mod/c"
	"git.parallelcoin.io/dev/jorm/pkg/gelook"
)

type Jorm struct {
	Window   *app.Window
	Context  *layout.Context
	Theme    *gelook.DuoUItheme
	Coins    c.Coins
	Selected c.Coin
}

func NewJorm() *Jorm {
	gofont.Register()
	w := app.NewWindow(
		app.Size(unit.Dp(1000), unit.Dp(800)),
		app.Title("ParallelCoin"),
	)
	return &Jorm{
		Window: w,
		Theme:  gelook.NewDuoUItheme(),
	}
}
