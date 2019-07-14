package main

import "fmt"

//https://tour.go-zh.org/moretypes/26
// 返回一个“返回int的函数”
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		temp := a
		a, b = b, (a + b) //重新赋值 重要！
		// fmt.Println("a,b: ", a, b)
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
