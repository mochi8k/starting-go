package main

import (
	"fmt"
)

func main() {
	// Unicodeコードポイントを表す特殊な整数型(int32のalias)
	var r1 rune = '松'
	fmt.Println(r1) // "26494"

	r2 := '松'
	fmt.Println(r2) // "26494"
}
