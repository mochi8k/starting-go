package main

import (
	f "fmt"
)

/*
アドレス演算子「&」
演算子&を使って、任意の型からそのポインタ型を生成することができる.
*/

func main() {
	var i int
	p := &i
	f.Printf("%T\n", p) // => "*int"

	pp := &p
	f.Printf("%T\n", pp) // => "**int"

	dereference()
}

func dereference() {
	var i int
	p := &i
	i = 5

	/* 変数の前に * を置くことで、デリファレンス可能. */
	f.Println(*p) // => "5"

	*p = 10
	f.Println(i) // => "10"
}
