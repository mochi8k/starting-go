package main

import (
	"fmt"
)

/*
  関数定義の基本
*/
func plus(x, y int) int {
	return x + y
}

/*
  引数を複数返す場合
*/
func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func main() {
	fmt.Println(plus(1, 5)) // => 6

	q, r := div(19, 7)
	fmt.Println(q, r) // => 2 5

	/*
	   戻り値の破棄
	*/
	// _, _ := div(19, 7) // 代入するとコンパイルエラー
	_, _ = div(19, 7)
	div(19, 7)
}
