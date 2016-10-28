package main

import (
	"fmt"
)

func main() {
	/*
	   interface{}とnil.
	   interface{}はObjectのようなもので、全ての型との互換を持つ.
	*/
	var x interface{}
	fmt.Println(x) // => "<nil>"

	x = 1
	x = 3.14
	x = '山'
	x = "文字列"
	x = [...]uint8{1, 2, 3, 4, 5}
	fmt.Println(x) // => [1, 2, 3, 4, 5]

	var n1, n2 interface{} = 1, 2

	// 一度interface{}型の変数に格納されると、整数動詞の加算など、データ型特有の演算ができなくなる.
	// interface{}はあくまでもすべての型の値を汎用的に表す手段であり、演算の対象にはできない.
	// n3 := n1 + n2 // エラー
	fmt.Println(n1, n2)
}
