package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Walk 遍历 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	a1 := make([]int, 0)
	a2 := make([]int, 0)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v := range ch1 {
		a1 = append(a1, v)
	}
	for v := range ch2 {
		a2 = append(a2, v)
	}
	if len(a1) != len(a2) {
		return false
	}

	// 是否应当Sort？
	// 题目要求的是：检测 t1 和 t2 是否存储了相同的值
	sort.Ints(a1)
	sort.Ints(a2)
	for i, v := range a1 {
		if v != a2[i] {
			return false
		}
	}
	return true
}

func main() {

	// Test Walk
	ch := make(chan int)
	t := tree.New(1)
	fmt.Println("tree:", t)
	go Walk(t, ch)
	for v := range ch {
		fmt.Println("Got value:", v)
	}

	// Test Same
	t2 := tree.New(1)
	fmt.Println("Check same:", Same(t, t2))
}
