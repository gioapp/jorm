package jorm

import (
	"gioui.org/io/system"
	"gioui.org/layout"
	"git.parallelcoin.io/dev/jorm/pkg/gel"
)

var (
	coinsPanelElement = gel.NewPanel()
)

func (j *Jorm) Gui() {
	go func() {
		j.Context = layout.NewContext(j.Window.Queue())
		for e := range j.Window.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				j.Context.Reset(e.Config, e.Size)
				layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(j.Context,
					layout.Flexed(0.5, j.CoinsList()),
					layout.Flexed(0.5, j.SelectedCoin()),
				)
				e.Frame(j.Context.Ops)
			}
		}
	}()
}