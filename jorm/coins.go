package jorm

import (
	"encoding/base64"
	"fmt"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"git.parallelcoin.io/dev/jorm/jdb"
	"golang.org/x/image/draw"
	"image"
	"strings"
)

func (j *Jorm) CoinsList() func() {
	return func() {
		coinsPanelElement.PanelObject = j.Coins.C
		coinsPanelElement.PanelObjectsNumber = len(j.Coins.C)
		coinsPanel := j.Theme.DuoUIpanel()
		coinsPanel.ScrollBar = j.Theme.ScrollBar()
		coinsPanel.Layout(j.Context, coinsPanelElement, func(i int, in interface{}) {
			coin := j.Coins.C[i]
			width := j.Context.Constraints.Width.Max
			layout.Flex{Axis: layout.Vertical}.Layout(j.Context,
				layout.Rigid(func() {
					j.Context.Constraints.Width.Min = width
					layout.Flex{
						Spacing: layout.SpaceBetween,
					}.Layout(j.Context,
						layout.Rigid(func() {
							//var logo image.Image
							//go func() {
							l := jdb.Read("data/"+coin.Slug, "logo")
							logos := l.(map[string]interface{})

							reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(logos["img16"].(string)))
							logo, _, err := image.Decode(reader)
							if err != nil {
								//log.Fatal(err)
							}
							if logo != nil {

								//bounds := m.Bounds()
								//}()
								fmt.Println("logoSloug", coin.Slug)
								fmt.Println("logoSloug", coin.Slug)
								imgOp := paint.ImageOp{}

								sz := 16
								if imgOp.Size().X != sz {
									imgRender := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz, Y: sz}})
									draw.ApproxBiLinear.Scale(imgRender,
										imgRender.Bounds(),
										logo,
										logo.Bounds(),
										draw.Src, nil)
									imgOp = paint.NewImageOp(imgRender)
								}
								imgRender := j.Theme.Image(imgOp)
								imgRender.Scale = float32(sz) / float32(j.Context.Px(unit.Dp(float32(sz))))
								imgRender.Layout(j.Context)
							}
						}),
						layout.Rigid(func() {
							j.Theme.H6(coin.Name).Layout(j.Context)
						}),
						layout.Rigid(func() {
							j.Theme.H6(coin.Ticker).Layout(j.Context)
						}),
					)
				}),
				layout.Rigid(j.Theme.DuoUIline(j.Context, 1, 0, 1, j.Theme.Colors["Gray"])),
			)
			//}
		})

		//j.Theme.H6("aaaa").Layout(j.Context)
	}
}
