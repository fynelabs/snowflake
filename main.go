//go:generate fyne bundle -o data.go data

package main

import (
	"fmt"
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeUI(s *snow) fyne.CanvasObject {
	merry := canvas.NewText("Merry Christmas", color.White)
	merry.Alignment = fyne.TextAlignCenter
	merry.TextSize = 42
	to := widget.NewLabel("To all our customers, investors\nand the whole Fyne community")
	to.Alignment = fyne.TextAlignCenter
	texts := container.NewVBox(merry,
		container.NewHBox(widget.NewLabel("From the Fyne Labs team"),
			layout.NewSpacer(),
			widget.NewButton("Sponsor", func() {
				u, _ := url.Parse("https://fyne.io/sponsor/")
				_ = fyne.CurrentApp().OpenURL(u)
			})))

	flake := newSnowFlake(2)
	label := widget.NewLabel(fmt.Sprintf("%d", flake.Count))
	down := widget.NewButtonWithIcon("", theme.MoveDownIcon(), func() {
		if flake.Count <= 0 {
			return
		}

		flake.Count--
		label.SetText(fmt.Sprintf("%d", flake.Count))
		flake.Refresh()
	})
	down.Importance = widget.LowImportance
	up := widget.NewButtonWithIcon("", theme.MoveUpIcon(), func() {
		if flake.Count >= 5 {
			return
		}

		flake.Count++
		label.SetText(fmt.Sprintf("%d", flake.Count))
		flake.Refresh()
	})
	up.Importance = widget.LowImportance
	start := widget.NewButton("Snow", func() {
		s.snow()
	})
	start.Importance = widget.LowImportance

	return container.NewBorder(to, texts, nil, nil, flake,
		container.NewCenter(container.NewVBox(container.NewHBox(
			down, label, up),
			start)))
}

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.FromLegacy(&cardTheme{}))
	w := a.NewWindow("Snowflake")

	s := newSnowLayer()
	w.SetPadded(false)
	w.SetContent(container.NewMax(
		canvas.NewRectangle(&color.NRGBA{R: 0, G: 0, B: 0x4d, A: 0xff}),
		container.NewPadded(container.NewPadded(makeUI(s))),
		s))
	w.Resize(fyne.NewSize(360, 520))
	w.ShowAndRun()
}
