package main

import (
	"fmt"
	"math"
)

//https://tour.go-zh.org/flowcontrol/8
func Sqrt(x float64) float64 {
	z := float64(1)    //声明z
	temp := float64(1) //保存临时变量
	for {
		z -= (z*z - x) / (2 * z) //计算最新的z值
		fmt.Println(z)
		if math.Abs(z-temp) <= 0.000000000000001 {
			break //当值停止改变（或改变非常小）的时候退出循环
		} else {
			temp = z
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
