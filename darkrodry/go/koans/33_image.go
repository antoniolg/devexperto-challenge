package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{
	height int
	width int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(x), 255}
}

func main() {
	m := Image{120, 255}
	pic.ShowImage(m)
}
