package main

import (
	"errors"
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

/*
  エラーを返す
*/
func errFn() (int, error) {
	return 1, errors.New("fact: domain error")
}

/*
  戻り値を表す変数を持つ
*/
func fn1() (a int) {
	return // a == 0
}

func fn2() int {
	var a int
	return a // a == 0
}

func fn3() (x, y int) {
	y = 5
	return // x == 0, y == 5
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

	/*
	   エラー処理
	*/
	result, err := errFn()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
