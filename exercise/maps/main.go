package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

//https://tour.go-zh.org/moretypes/23
func WordCount(s string) map[string]int {
	m := make(map[string]int) //创建map
	d := strings.Fields(s)    //将带空格的s字符串变成一个字符串数组 https://go-zh.org/pkg/strings/#Fields
	for _, v := range d {
		m[v] = m[v] + 1 //相同单词出现一次就加1
	}
	return m

	//return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
