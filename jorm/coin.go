package jorm

import (
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/gioapp/gelook"
	"golang.org/x/image/draw"
	"image"
)

var (
	coinDataList = &layout.List{
		Axis: layout.Vertical,
	}
)

func (j *Jorm) SelectedCoin() func() {
	return func() {
		j.Theme.DuoUIcontainer(16, j.Theme.Colors["DarkGrayI"]).Layout(j.Context, layout.N, func() {
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(j.Context,
				layout.Rigid(j.coinHeader()),
				layout.Flexed(1, j.coinBody()),
				layout.Rigid(j.coinFooter()),
			)
		})
	}
}

func InfoRow(gtx *layout.Context, th *gelook.DuoUItheme, width int, label, data string) func() {
	return func() {
		gtx.Constraints.Width.Min = width - 16
		layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func() {
				layout.Flex{
					Spacing: layout.SpaceBetween,
				}.Layout(gtx,
					layout.Rigid(func() {
						th.H6(label).Layout(gtx)
					}),
					layout.Rigid(func() {
						th.Body1(data).Layout(gtx)
					}),
				)
			}),
			layout.Rigid(th.DuoUIline(gtx, 2, 0, 1, th.Colors["Primary"])),
		)
	}
}

func (j *Jorm) coinData(width int) []func() {
	return []func(){
		InfoRow(j.Context, j.Theme, width, "Description", j.Selected.Data.Description),
		InfoRow(j.Context, j.Theme, width, "Algo", j.Selected.Data.Algo),
		InfoRow(j.Context, j.Theme, width, "Github", j.Selected.Data.Github),
		InfoRow(j.Context, j.Theme, width, "BlockRewardReduction", j.Selected.Data.BlockRewardReduction),
		InfoRow(j.Context, j.Theme, width, "ProofType", j.Selected.Data.ProofType),
		InfoRow(j.Context, j.Theme, width, "DifficultyAdjustment", j.Selected.Data.DifficultyAdjustment),
		InfoRow(j.Context, j.Theme, width, "Reddit", j.Selected.Data.Reddit),
		InfoRow(j.Context, j.Theme, width, "WhitePaper", j.Selected.Data.WhitePaper),
		InfoRow(j.Context, j.Theme, width, "WebSite", j.Selected.Data.WebSite),
		InfoRow(j.Context, j.Theme, width, "Twitter", j.Selected.Data.Twitter),
		InfoRow(j.Context, j.Theme, width, "TotalCoinSupply", j.Selected.Data.TotalCoinSupply),
		InfoRow(j.Context, j.Theme, width, "StartDate", j.Selected.Data.StartDate),
	}
}

func (j *Jorm) coinHeader() func() {
	return func() {
		j.Theme.DuoUIcontainer(8, "ffffffff").Layout(j.Context, layout.Center, func() {
			layout.Flex{
				Alignment: layout.Middle,
				Spacing:   layout.SpaceBetween,
			}.Layout(j.Context,
				layout.Rigid(func() {
					j.Theme.DuoUIcontainer(0, "ffffffff").Layout(j.Context, layout.Center, func() {
						if j.Selected.Logo != nil {
							imgOp := paint.ImageOp{}
							sz := 128
							if imgOp.Size().X != sz {
								imgRender := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz, Y: sz}})
								draw.ApproxBiLinear.Scale(imgRender,
									imgRender.Bounds(),
									j.Selected.LogoBig,
									j.Selected.LogoBig.Bounds(),
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
						j.Theme.H4(j.Selected.Name).Layout(j.Context)
					})
				}),
				layout.Rigid(func() {
					j.Theme.H3(j.Selected.Ticker).Layout(j.Context)
				}),
			)
		})

	}
}
func (j *Jorm) coinBody() func() {
	return func() {
		body := j.Theme.DuoUIcontainer(0, "")
		body.PaddingTop = 16
		body.PaddingBottom = 16
		body.Layout(j.Context, layout.N, func() {
			j.Theme.DuoUIcontainer(0, "").Layout(j.Context, layout.Center, func() {
				width := j.Context.Constraints.Width.Max
				layout.Flex{
					Alignment: layout.Middle,
					Spacing:   layout.SpaceBetween,
				}.Layout(j.Context,
					layout.Rigid(func() {
						j.Context.Constraints.Width.Min = width

						j.Theme.DuoUIcontainer(8, "ffffffff").Layout(j.Context, layout.Center, func() {
							coinDataList.Layout(j.Context, len(j.coinData(width)), func(i int) {
								j.Theme.DuoUIcontainer(0, "ffffffff").Layout(j.Context, layout.Center, func() { j.coinData(width)[i]() })
							})
						})
					}),
					layout.Flexed(1, func() {
						//title := j.Theme.DuoUIcontainer(0, "")
						//title.PaddingLeft = 8
						//title.Layout(j.Context, layout.W, func() {
						//	j.Theme.H4(j.Selected.Name).Layout(j.Context)
						//})
					}),
					layout.Rigid(func() {
						//j.Theme.H3(j.Selected.Ticker).Layout(j.Context)
					}),
				)
			})
		})
	}
}

func (j *Jorm) coinFooter() func() {
	return func() {
		j.Theme.DuoUIcontainer(8, "ffffffff").Layout(j.Context, layout.Center, func() {
			layout.Flex{
				Alignment: layout.Middle,
				Spacing:   layout.SpaceBetween,
			}.Layout(j.Context,
				layout.Rigid(func() {
					j.Theme.DuoUIcontainer(0, "ffffffff").Layout(j.Context, layout.Center, func() {
						if j.Selected.Logo != nil {
							imgOp := paint.ImageOp{}
							sz := 128
							if imgOp.Size().X != sz {
								imgRender := image.NewRGBA(image.Rectangle{Max: image.Point{X: sz, Y: sz}})
								draw.ApproxBiLinear.Scale(imgRender,
									imgRender.Bounds(),
									j.Selected.LogoBig,
									j.Selected.LogoBig.Bounds(),
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
						j.Theme.H4(j.Selected.Name).Layout(j.Context)
					})
				}),
				layout.Rigid(func() {
					j.Theme.H3(j.Selected.Ticker).Layout(j.Context)
				}),
			)
		})
	}
}
