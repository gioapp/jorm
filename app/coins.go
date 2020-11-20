package jorm

import (
	l "gioui.org/layout"
)

func (j *Jorm) CoinsList(gtx l.Context) l.Dimensions {
	gtx.Constraints.Max.X = 300

	//le := func(gtx l.Context, index int) l.Dimensions {
	//	return createWalletLayoutList[index](gtx)
	//}

	return j.lists["coins"].
		End().
		Vertical().
		Length(len(j.Coins.C)).
		ScrollToEnd().
		DisableScroll(false).
		ListElement(func(gtx l.Context, index int) l.Dimensions {
			c := j.Coins.C[index]
			//if c.Link.Clicked() {
			//	j.Selected = *c.SelectCoin()
			//}
			background := "PanelBg"
			//color := "PanelText"
			//c.Link.Events
			img := func(gtx l.Context) l.Dimensions { return l.Dimensions{} }
			return j.th.ButtonLayout(c.Link).
				CornerRadius(0).
				Embed(j.th.Inset(0.066,
					j.th.Flex().
						Rigid(
							func(gtx l.Context) l.Dimensions {
								if c.Logo != nil {
									img = j.th.Image().Scale(2).Fn
								}
								return img(gtx)
							}).
						Flexed(1,
							j.th.H6(c.Ticker).Fn,
						).
						Rigid(
							j.th.H6(c.Name).Fn,
						).Fn).Fn).
				Background(background).
				SetClick(
					func() {
						//if j.MenuOpen {
						//	j.MenuOpen = false
						//}
						//j.ActivePage(name)
					}).Fn(gtx)
		}).Fn(gtx)

}
