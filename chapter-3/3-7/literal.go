package main

import (
	"fmt"
)

func main() {
	a := 123        // 10進数
	b := 0755       // 8進数
	c := 0x0719BEEF // 16進数

	fmt.Println(a, b, c)
}
