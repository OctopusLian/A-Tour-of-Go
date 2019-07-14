package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

//实现Image包中演示模式的方法
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

//实现Image包中生成图片边界的方法
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

//实现Image包中生成图像某个点的方法
func (i Image) At(x, y int) color.Color {
	return color.RGBA{
		uint8(x),
		uint8(y),
		uint8(255),
		uint8(255),
	}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
