package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	pixels [][]color.Color
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0, len(i.pixels[0]), len(i.pixels))
}

func (i Image) At(x, y int) color.Color {
	return i.pixels[y][x]
}

func Pic(dx, dy int) [][]color.Color {
	sl := make([][]color.Color, dy)
	for i := range sl {
		sl[i] = make([]color.Color, dx)
		for j := 0; j < dx; j++ {
			sl[i][j] = color.RGBA{uint8(i * j), uint8(i + j), 255, 255}
		}
	}
	return sl
}

func main() {
	m := Image{Pic(50,100)}
	pic.ShowImage(m)
}
