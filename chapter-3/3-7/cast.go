package main

import (
	"fmt"
)

func main() {
	// 整数リテラルを使って暗黙的に定義した変数はint型となる
	n := 17

	un := uint(17)
	b := byte(n)
	i64 := int64(n)
	u32 := uint32(n)

	fmt.Println(n, un, b, i64, u32)
}
