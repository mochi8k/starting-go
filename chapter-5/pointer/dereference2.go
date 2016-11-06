package main

import (
	f "fmt"
)

func pow(p *[3]int) {
	i := 0
	for i < 3 {
		/* 各要素を累乗 */
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}

func main() {
	p := &[3]int{1, 2, 3}
	pow(p)
	f.Println(*p) // => "[1, 4, 9]"
}
