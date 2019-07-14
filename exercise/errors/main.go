package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) //当传入的数小于0时，返回0和错误信息
	}

	return math.Sqrt(x), nil //当传入的数大于等于0时，返回结果和nil
	// return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

/*
Output
1.4142135623730951 <nil>
0 cannot Sqrt negative number: -2
*/
