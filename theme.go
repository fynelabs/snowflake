package main

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type cardTheme struct{}

func (c cardTheme) BackgroundColor() color.Color {
	return &color.NRGBA{R: 0, G: 0, B: 0x4d, A: 0xff}
}

func (c cardTheme) ButtonColor() color.Color {
	return color.Transparent
}

func (c cardTheme) DisabledButtonColor() color.Color {
	return color.White
}

func (c cardTheme) HyperlinkColor() color.Color {
	return color.White
}

func (c cardTheme) TextColor() color.Color {
	return color.White
}

func (c cardTheme) DisabledTextColor() color.Color {
	return color.Black
}

func (c cardTheme) IconColor() color.Color {
	return color.White
}

func (c cardTheme) DisabledIconColor() color.Color {
	return color.Black
}

func (c cardTheme) PlaceHolderColor() color.Color {
	return color.White
}

func (c cardTheme) PrimaryColor() color.Color {
	return color.White
}

func (c cardTheme) HoverColor() color.Color {
	return &color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
}

func (c cardTheme) FocusColor() color.Color {
	return color.White
}

func (c cardTheme) ScrollBarColor() color.Color {
	return color.Black
}

func (c cardTheme) ShadowColor() color.Color {
	return color.White
}

func (c cardTheme) TextSize() int {
	return 18
}

func (c cardTheme) TextFont() fyne.Resource {
	return resourcePacificoRegularTtf
}

func (c cardTheme) TextBoldFont() fyne.Resource {
	return resourcePacificoRegularTtf
}

func (c cardTheme) TextItalicFont() fyne.Resource {
	return resourcePacificoRegularTtf
}

func (c cardTheme) TextBoldItalicFont() fyne.Resource {
	return resourcePacificoRegularTtf
}

func (c cardTheme) TextMonospaceFont() fyne.Resource {
	return resourcePacificoRegularTtf
}

func (c cardTheme) Padding() int {
	return 4
}

func (c cardTheme) IconInlineSize() int {
	return 24
}

func (c cardTheme) ScrollBarSize() int {
	return 18
}

func (c cardTheme) ScrollBarSmallSize() int {
	return 4
}

type overlayTheme struct{
	cardTheme
}

func (c overlayTheme) BackgroundColor() color.Color {
	return color.Transparent
}
