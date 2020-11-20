package jorm

import (
	"gioui.org/app"
	l "gioui.org/layout"
	"github.com/gioapp/jorm/mod/c"
	"github.com/p9c/pod/pkg/gui/fonts/p9fonts"
	"github.com/p9c/pod/pkg/gui/p9"
)

type Jorm struct {
	Window                                             *app.Window
	Coins                                              c.Coins
	Selected                                           c.Coin
	th                                                 *p9.Theme
	button0, button1, button2, iconbutton, iconbutton1 *p9.Clickable
	boolButton1, boolButton2                           *p9.Bool
	quit                                               chan struct{}
	progress                                           int
	slider                                             *p9.Float
	lineEditor, areaEditor                             *p9.Editor
	radio                                              *p9.Enum
	lists                                              map[string]*p9.List
}

func NewJorm() *Jorm {
	quit := make(chan struct{})
	th := p9.NewTheme(p9fonts.Collection(), quit)
	return &Jorm{
		th:      th,
		button0: th.Clickable(),
		button1: th.Clickable(),
		button2: th.Clickable().SetClick(func() {
			Info("clicked default style button")
		}),
		boolButton1: th.Bool(false),
		boolButton2: th.Bool(false),
		iconbutton:  th.Clickable(),
		iconbutton1: th.Clickable(),
		quit:        make(chan struct{}),
		progress:    0,
		slider: th.Float().SetHook(func(fl float32) {
			Debug("float now at value", fl)
		}),
		lineEditor: th.Editor().SingleLine().Submit(true),
		areaEditor: th.Editor().SingleLine().Submit(false),
		radio: th.Enum().SetOnChange(func(value string) {
			Debug("changed radio button to", value)
		}),
		lists: map[string]*p9.List{
			"coins": th.List(),
			"coin":  th.List(),
		},
	}
}

func (j *Jorm) Jorm(gtx l.Context) l.Dimensions {
	j.progress++
	if j.progress == 100 {
		j.progress = 0
	}
	th := j.th
	return th.Flex().Flexed(1, th.Flex().Rigid(th.Flex().Flexed(0.5, th.Fill("PanelBg", th.Inset(0.25, j.CoinsList).Fn).Fn).Flexed(0.5, th.Fill("DocBg", th.Inset(0.25, j.SelectedCoin()).Fn).Fn).Fn).Fn).Fn(gtx)
}
