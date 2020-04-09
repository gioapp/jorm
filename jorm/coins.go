package jorm

import (
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"golang.org/x/image/draw"
	"image"
)

func (j *Jorm) CoinsList() func() {
	return func() {
		j.Context.Constraints.Width.Max = 300
		coinsPanelElement.PanelObject = j.Coins.C
		coinsPanelElement.PanelObjectsNumber = len(j.Coins.C)
		coinsPanel := j.Theme.DuoUIpanel()
		coinsPanel.ScrollBar = j.Theme.ScrollBar(0)
		coinsPanel.Layout(j.Context, coinsPanelElement, func(i int, in interface{}) {
			coin := j.Coins.C[i]
			layout.Flex{Axis: layout.Vertical}.Layout(j.Context,
				layout.Rigid(func() {
					if coin.Link.Clicked(j.Context) {
						j.Selected = *coin.SelectCoin()
					}
					j.Theme.DuoUIbutton("", "", "", "", "", "", "", "", 0, 0, 0, 0, 0, 0, 0, 0).InsideLayout(j.Context, coin.Link, func() {
						layout.Flex{
							Alignment: layout.Middle,
							Spacing:   layout.SpaceBetween,
						}.Layout(j.Context,
							layout.Rigid(func() {
								j.Theme.DuoUIcontainer(0, "ffffffff").Layout(j.Context, layout.Center, func() {
									if coin.Logo != nil {
										imgOp := paint.ImageOp{}
										sz := 32
										if imgOp.Size().X != sz {
											imgRender := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz, Y: sz}})
											draw.ApproxBiLinear.Scale(imgRender,
												imgRender.Bounds(),
												coin.Logo,
												coin.Logo.Bounds(),
												draw.Src, nil)
											imgOp = paint.NewImageOp(imgRender)
										}
										imgRender := j.Theme.Image(imgOp)
										imgRender.Scale = float32(sz) / float32(j.Context.Px(unit.Dp(float32(sz))))
										imgRender.Layout(j.Context)
									}
								})
							}),
							layout.Flexed(1, func() {
								title := j.Theme.DuoUIcontainer(0, "")
								title.PaddingLeft = 8
								title.Layout(j.Context, layout.W, func() {
									j.Theme.H6(coin.Name).Layout(j.Context)
								})
							}),
							layout.Rigid(func() {
								title := j.Theme.DuoUIcontainer(0, "")
								title.PaddingRight = 8
								title.Layout(j.Context, layout.W, func() {
									j.Theme.H6(coin.Ticker).Layout(j.Context)
								})
							}),
						)
					})
				}),
				layout.Rigid(j.Theme.DuoUIline(j.Context, 1, 0, 1, j.Theme.Colors["LightGrayIII"])),
			)
			//}
		})

		//j.Theme.H6("aaaa").Layout(j.Context)
	}
}
