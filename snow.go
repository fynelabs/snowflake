package main

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var pix image.Image

func init() {
	pix = captureSnowflake(flakeSizeLarge.Add(flakeSizeLarge), false)
}

type snow struct {
	widget.BaseWidget
	flakes []fyne.CanvasObject
	bg     *canvas.Rectangle
}

func newSnowLayer() *snow {
	s := &snow{}
	s.ExtendBaseWidget(s)

	go s.animate()
	return s
}

func (s *snow) animate() {
	for {
		<-time.After(time.Millisecond * 16)
		fyne.Do(func() {
			flakes := make([]fyne.CanvasObject, len(s.flakes))
			remove := 0
			for i, f := range s.flakes {
				if f == nil {
					continue
				}
				if f.Size().Height > 40 {
					f.Move(f.Position().Add(fyne.NewPos(0, 1)))
				} else if f.Size().Height > 20 {
					f.Move(f.Position().Add(fyne.NewPos(0, 0.5)))
				} else {
					f.Move(f.Position().Add(fyne.NewPos(0, 0.25)))
				}

				if f.Position().Y < s.Size().Height+36 { // mobile overflow approximation
					flakes[i-remove] = f
				} else {
					remove++
				}
			}
			s.flakes = flakes[:len(flakes)-remove]
		})
	}
}

func (s *snow) snow() {
	space := int(s.Size().Width - flakeSizeLarge.Width)
	count := space / 50

	go func() {
		fyne.Do(func() {
			flakes := make([]fyne.CanvasObject, count)
			for i := 0; i < count; i++ {

				x := rand.Intn(space)
				y := rand.Intn(80)

				f := canvas.NewImageFromImage(pix)
				f.ScaleMode = canvas.ImageScaleFastest
				switch rand.Intn(6) {
				case 5:
					f.Resize(flakeSizeLarge)
				case 3, 4:
					f.Resize(flakeSizeMid)
				default:
					f.Resize(flakeSizeSmall)
				}
				f.Move(fyne.NewPos(float32(x), float32(-y)-f.Size().Height*2))

				flakes[i] = f

				s.flakes = append(s.flakes, flakes...)
			}
		})
	}()

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
