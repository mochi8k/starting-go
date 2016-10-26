package main

import (
	"fmt"
	"math"
)

func main() {
	// b2 := byte(256) コンパイルエラー
	n2 := 256

	// オーバーフローが発生した場合、ラップアラウンドさせる.
	b2 := byte(n2)

	fmt.Println(b2)

	fmt.Println(isWrapAround(42, 32))        // false
	fmt.Println(isWrapAround(3, 4294967295)) // true
}

func isWrapAround(a, b uint32) bool {
	return (math.MaxUint32 - a) < b
}
