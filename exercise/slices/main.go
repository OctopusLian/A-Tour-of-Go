package main

import (
	// "fmt"

	"golang.org/x/tour/pic"
)

//https://tour.go-zh.org/moretypes/18
func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy) //创建外层切片
	// fmt.Println("a: ", a)
	for x := range a {
		b := make([]uint8, dx) //创建里层切片
		for y := range b {
			/*
				图像的选择由你来定。几个有趣的函数包括 (x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)。
			*/
			b[y] = uint8(x*y - 1) //给里层切片的每一个元素赋值
			// fmt.Println("b[y]: ", b[y])
		}
		a[x] = b
		// fmt.Println("a[x]: ", a[x])
	}
	return a
}

func main() {
	pic.Show(Pic)
}
