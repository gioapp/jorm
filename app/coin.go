package jorm

import (
	l "gioui.org/layout"
	"gioui.org/op/paint"
	"github.com/p9c/pod/pkg/gui/p9"
	"golang.org/x/image/draw"
	"image"
)

func (j *Jorm) SelectedCoin() func(gtx l.Context) l.Dimensions {
	return j.th.Flex().Rigid(j.coinHeader).Flexed(1, j.coinBody).Rigid(j.coinHeader).Rigid(j.coinFooter).Fn
}

func InfoRow(th *p9.Theme, width int, label, data string) func(gtx l.Context) l.Dimensions {
	return func(gtx l.Context) l.Dimensions {
		gtx.Constraints.Min.X = width - 16
		return th.Flex().Rigid(th.Flex().Rigid(th.H6(label).Fn).Rigid(th.Body1(data).Fn).Fn).Fn(gtx)
	}
}

func (j *Jorm) coinData(width int) []func(gtx l.Context) l.Dimensions {
	return []func(gtx l.Context) l.Dimensions{
		InfoRow(j.th, width, "Description", j.Selected.Data.Description),
		InfoRow(j.th, width, "Algo", j.Selected.Data.Algo),
		InfoRow(j.th, width, "Github", j.Selected.Data.Github),
		InfoRow(j.th, width, "BlockRewardReduction", j.Selected.Data.BlockRewardReduction),
		InfoRow(j.th, width, "ProofType", j.Selected.Data.ProofType),
		InfoRow(j.th, width, "DifficultyAdjustment", j.Selected.Data.DifficultyAdjustment),
		InfoRow(j.th, width, "Reddit", j.Selected.Data.Reddit),
		InfoRow(j.th, width, "WhitePaper", j.Selected.Data.WhitePaper),
		InfoRow(j.th, width, "WebSite", j.Selected.Data.WebSite),
		InfoRow(j.th, width, "Twitter", j.Selected.Data.Twitter),
		InfoRow(j.th, width, "TotalCoinSupply", j.Selected.Data.TotalCoinSupply),
		InfoRow(j.th, width, "StartDate", j.Selected.Data.StartDate),
	}
}

func (j *Jorm) coinHeader(gtx l.Context) l.Dimensions {
	return j.th.Flex().Rigid(
		func(gtx l.Context) l.Dimensions {
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
				//imgRender := j.th.Image(imgOp)
				//imgRender.Scale = float32(sz) / float32(gtx.Px(unit.Dp(float32(sz))))
				//imgRender.Layout(gtx)
			}
			//})
			return l.Dimensions{}
		}).Flexed(1,
		//title := j.th.DuoUIcontainer(0, "")
		//title.PaddingLeft = 8
		//title.Layout(gtx, l.W, func() {
		j.th.H4(j.Selected.Name).Fn,
	).
		Rigid(j.th.H3(j.Selected.Ticker).Fn).Fn(gtx)
}

func (j *Jorm) coinBody(gtx l.Context) l.Dimensions {
	width := gtx.Constraints.Max.X
	return j.th.Flex().Rigid(func(gtx l.Context) l.Dimensions {
		gtx.Constraints.Min.X = width

		return j.lists["coin"].
			End().
			Vertical().
			Length(len(j.coinData(width))).
			ScrollToEnd().
			DisableScroll(false).
			ListElement(func(gtx l.Context, index int) l.Dimensions {
				return j.coinData(width)[index](gtx)
			}).Fn(gtx)
		// Slice(gtx, widgets...)(gtx)
	}).Flexed(1, func(gtx l.Context) l.Dimensions {
		//title := j.th.DuoUIcontainer(0, "")
		//title.PaddingLeft = 8
		//title.Layout(gtx, l.W, func() {
		//	j.th.H4(j.Selected.Name).Layout(gtx)
		//})
		return l.Dimensions{}
	}).Rigid(func(gtx l.Context) l.Dimensions {
		//j.th.H3(j.Selected.Ticker).Layout(gtx)
		return l.Dimensions{}
	}).Fn(gtx)
}

func (j *Jorm) coinFooter(gtx l.Context) l.Dimensions {
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
		//imgRender := j.th.Image(imgOp)
		//imgRender.Scale = float32(sz) / float32(gtx.Px(unit.Dp(float32(sz))))
		//imgRender.Layout(gtx)
	}

	return j.th.Flex().Flexed(1, func(gtx l.Context) l.Dimensions {
		//title.PaddingLeft = 8
		//title.Layout(gtx, l.W, func(gtx l.Context) l.Dimensions {
		//	j.th.H4(j.Selected.Name).Layout(gtx)
		//})
		return l.Dimensions{}

	}).Rigid(j.th.H3(j.Selected.Ticker).Fn).Fn(gtx)
}
