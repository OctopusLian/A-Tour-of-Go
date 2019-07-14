package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//再封装io.Reader
type rot13Reader struct {
	r io.Reader
}

//字母交换，遵循rot13转换规则
func rot13(out byte) byte {
	switch {
	case out >= 'A' && out <= 'M' || out >= 'a' && out <= 'm':
		out += 13
	case out >= 'N' && out <= 'Z' || out >= 'n' && out <= 'z':
		out -= 13
	}

	return out
}

//重写Read方法
func (rot13r rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13r.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}

	fmt.Println("result is ", string(b))

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
