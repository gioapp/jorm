package jorm

import (
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/gioapp/gel"
)

var (
	coinsPanelElement = gel.NewPanel()
)

func (j *Jorm) Gui() {
	j.Context = layout.NewContext(j.Window.Queue())
	for e := range j.Window.Events() {
		if e, ok := e.(system.FrameEvent); ok {
			j.Context.Reset(e.Config, e.Size)
			layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(j.Context,
				layout.Rigid(j.CoinsList()),
				layout.Flexed(1, j.SelectedCoin()),
			)
			e.Frame(j.Context.Ops)
		}
	}
}
