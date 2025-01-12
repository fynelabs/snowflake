package main

import (
	"image"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/tools/playground"
	"fyne.io/fyne/v2/widget"
)

var (
	flakeSizeSmall = fyne.Size{16, 16}
	flakeSizeMid   = fyne.Size{24, 24}
	flakeSizeLarge = fyne.Size{48, 48}
)

type snowflake struct {
	widget.BaseWidget
	Count int
}

func newSnowFlake(c int) *snowflake {
	s := &snowflake{Count: c}
	s.ExtendBaseWidget(s)
	return s
}

func (s *snowflake) CreateRenderer() fyne.WidgetRenderer {
	return &snowflakeRender{flake: s,
		side1: []*canvas.Line{},
		side2: []*canvas.Line{},
		side3: []*canvas.Line{}}
}

type snowflakeRender struct {
	flake *snowflake

	oldSize  fyne.Size
	oldCount int

	side1, side2, side3 []*canvas.Line
}

func (r *snowflakeRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *snowflakeRender) Destroy() {
}

func (r *snowflakeRender) Layout(s fyne.Size) {
	if r.flake.Count == r.oldCount && s.Subtract(r.oldSize).IsZero() {
		return
	}
	r.oldCount = r.flake.Count
	r.oldSize = s

	hh := float32(float64(s.Height/2) * math.Sqrt(3.0))
	len := fyne.Min(s.Width, hh)
	height := float32(float64(len/2) * math.Sqrt(3.0))
	off := fyne.NewPos(0, 0)
	if s.Width > hh {
		off.X = (s.Width - len) / 2
	} else {
		off.Y = (hh - len) / 2
	}

	top := off.Add(fyne.NewPos(len/2, 0))
	left := off.Add(fyne.NewPos(0, height))
	right := left.Add(fyne.NewPos(len, 0))

	r.side1 = koch(left, top, len, 300, r.flake.Count)
	r.side2 = koch(top, right, len, 60, r.flake.Count)
	r.side3 = koch(right, left, len, 180, r.flake.Count)
}

func (r *snowflakeRender) MinSize() fyne.Size {
	return fyne.NewSize(16, 16)
}

func (r *snowflakeRender) Objects() []fyne.CanvasObject {
	obj := make([]fyne.CanvasObject, len(r.side1)+len(r.side2)+len(r.side3))
	for i, l := range r.side1 {
		obj[i] = l
	}
	for i, l := range r.side2 {
		obj[len(r.side1)+i] = l
	}
	for i, l := range r.side3 {
		obj[len(r.side1)+len(r.side2)+i] = l
	}
	return obj
}

func (r *snowflakeRender) Refresh() {
	r.Layout(r.flake.Size())
}

func newLine(start, stop fyne.Position) *canvas.Line {
	l := canvas.NewLine(theme.TextColor())
	l.StrokeWidth = 2.5
	l.Position1 = start
	l.Position2 = stop
	return l
}

func koch(start, stop fyne.Position, len float32, angle, count int) []*canvas.Line {
	if count == 0 {
		return []*canvas.Line{newLine(start, stop)}
	}

	diff := fyne.NewPos((stop.X-start.X)/3, (stop.Y-start.Y)/3)
	bend1 := start.Add(diff)
	bend2 := stop.Subtract(diff)

	ang := angle - 60
	rad := float64(ang) * math.Pi / 180.0
	x2 := float32(math.Cos(rad) * float64(len/3))
	y2 := float32(math.Sin(rad) * float64(len/3))
	mid := bend1.Add(fyne.NewPos(x2, y2))

	lines := []*canvas.Line{}
	for _, line := range koch(start, bend1, len/3, angle, count-1) {
		lines = append(lines, line)
	}
	for _, line := range koch(bend1, mid, len/3, ang, count-1) {
		lines = append(lines, line)
	}
	for _, line := range koch(mid, bend2, len/3, ang+120, count-1) {
		lines = append(lines, line)
	}
	for _, line := range koch(bend2, stop, len/3, angle, count-1) {
		lines = append(lines, line)
	}
	return lines
}

func captureSnowflake(s fyne.Size, bg bool) image.Image {
	c := playground.NewSoftwareCanvas()
	c.SetPadded(false)
	fyne.CurrentApp().Settings().SetTheme(theme.FromLegacy(&overlayTheme{}))

	var content fyne.CanvasObject
	if bg {
		content = container.NewMax(
			canvas.NewRectangle(&color.NRGBA{R: 0, G: 0, B: 0x4d, A: 0xff}),
			container.NewPadded(newSnowFlake(2)))
	} else {
		content = container.NewPadded(newSnowFlake(2))
	}
	c.SetContent(content)
	c.Resize(s)
	c.SetScale(2)

	img := c.Capture()
	fyne.CurrentApp().Settings().SetTheme(theme.FromLegacy(&cardTheme{}))
	return img
}
