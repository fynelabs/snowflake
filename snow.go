package main

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
)

var pix image.Image

func init() {
	pix = captureSnowflake(flakeSizeLarge.Add(flakeSizeLarge), false)
}

type snow struct {
	widget.BaseWidget
	flakes []fyne.CanvasObject
}

func newSnowLayer() *snow {
	s := &snow{}
	s.ExtendBaseWidget(s)

	go s.animate()
	return s
}

func (s *snow) animate() {
	even := false
	for {
		<-time.After(time.Millisecond * 33)
		var flakes []fyne.CanvasObject
		for _, f := range s.flakes {
			if f.Size().Height > 40 {
				f.Move(f.Position().Add(fyne.NewPos(0, 2)))
			} else if f.Size().Height > 20 {
				f.Move(f.Position().Add(fyne.NewPos(0, 1)))
			} else if even { // half speed for small
				f.Move(f.Position().Add(fyne.NewPos(0, 1)))
			}

			if f.Position().Y < s.Size().Height + 36 { // mobile overflow approximation
				flakes = append(flakes, f)
			}
		}

		even = !even
		s.flakes = flakes
		canvas.Refresh(s)
	}
}

func (s *snow) snow() {
	space := s.Size().Width - flakeSizeLarge.Width
	count := space / 50
	for i := 0; i < count; i++ {

		x := rand.Intn(space)
		y := rand.Intn(80)

		f := canvas.NewImageFromImage(pix)
		switch rand.Intn(6) {
		case 5:
			f.Resize(flakeSizeLarge)
		case 3, 4:
			f.Resize(flakeSizeMid)
		default:
			f.Resize(flakeSizeSmall)
		}
		f.Move(fyne.NewPos(x, -y-f.Size().Height*2))

		s.flakes = append(s.flakes, f)
	}

	canvas.Refresh(s)
}

func (s *snow) CreateRenderer() fyne.WidgetRenderer {
	return &snowRender{s: s}
}

type snowRender struct {
	s *snow
}

func (s *snowRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (s *snowRender) Destroy() {
	// TODO stop all animations
}

func (s *snowRender) Layout(fyne.Size) {
	// no-op all the animations are manually placed
}

func (s *snowRender) MinSize() fyne.Size {
	return flakeSizeSmall
}

func (s *snowRender) Objects() []fyne.CanvasObject {
	return s.s.flakes
}

func (s *snowRender) Refresh() {
	// no-op
}
