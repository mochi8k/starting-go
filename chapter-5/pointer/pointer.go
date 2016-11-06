package main

import (
	f "fmt"
)

func main() {
	/* int型のポインタ */
	var p1 *int
	f.Println(p1 == nil) // true

	/* int型のポインタのポインタ型 */
	var p2 ***int
	f.Println(p2 == nil) // true

	/*
		参照型にもポインタ型を定義できるが、
		参照型はそもそもポインタを使った参照を含んでいるため、
		必要になる局面は、相当なレアケースに限られる.
	*/
	var (
		s  *[]string
		m  *map[int]rune
		ch *chan int
	)
	f.Println(s == nil)  // true
	f.Println(m == nil)  // true
	f.Println(ch == nil) // true
}
