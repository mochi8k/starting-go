package main

import (
	"fmt"
)

// 前に定義されている定数の値が、以降の定数となる.
const (
	X  = 1   // == 1
	Y        // == 1
	Z        // == 1
	S1 = "あ" // == "あ"
	S2       // == "あ"
)

// コンパイルエラー
// const (
// 	A
// 	B
// 	C
// )

// 定数値の式
const (
	A = 2
	B = 7
	C = A + B // == 9

	S3 = "今日"
	S4 = "晴れ"
	S5 = S3 + "は" + S4 // == "今日は晴れ"
)

// 定義順序は関係ない
const (
	F = D + E // == 9
	D = 2
	E = 7
)

// パッケージ定数
const ONE = 1

func one() (int, int) {
	// 関数内定数
	const TWO = 2
	return ONE, TWO
}

func main() {
	x, y := one()
	fmt.Println(x, y) // => 1 2
}
