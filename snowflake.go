package main

import (
	"image/color"
	"math"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
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

	side1, side2, side3 []*canvas.Line
}

func (r *snowflakeRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *snowflakeRender) Destroy() {
}

func (r *snowflakeRender) Layout(s fyne.Size) {
	hh := int(float64(s.Height/2) * math.Sqrt(3.0))
	len := fyne.Min(s.Width, hh)
	height := int(float64(len/2) * math.Sqrt(3.0))
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
	canvas.Refresh(r.flake)
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
	l.Position1 = start
	l.Position2 = stop
	return l
}

func koch(start, stop fyne.Position, len, angle, count int) []*canvas.Line {
	if count == 1 {
		return []*canvas.Line{newLine(start, stop)}
	}

	diff := fyne.NewPos((stop.X-start.X)/3, (stop.Y-start.Y)/3)
	bend1 := start.Add(diff)
	bend2 := stop.Subtract(diff)

	ang := angle - 60
	rad := float64(ang) * math.Pi / 180.0
	x2 := int(math.Cos(rad) * float64(len/3))
	y2 := int(math.Sin(rad) * float64(len/3))
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
