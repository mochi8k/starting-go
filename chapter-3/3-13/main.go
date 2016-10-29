package main

import (
	myFoo "./foo" // 別名をつけることが可能
	. "fmt"       // . でパッケージ名の省略が可能になる
)

const (
	Println = "hello" // fmt.Printlnと重複するためエラー
)

func main() {
	Println(myFoo.MAX)
	// Println(myFoo.internal_const) // コンパイルエラー

	Println(myFoo.FooFunc(5))
	// Println(myFoo.InternalFunc(5)) // コンパイルエラー
}
