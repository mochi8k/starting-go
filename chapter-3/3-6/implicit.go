package main

import "fmt"

func one() int {
	return 1
}

func main() {
	i := 1      // int
	b := true   // bool
	f := 3.14   // float64
	s := "abc"  // string
	i1 := one() // 関数の戻り値の型

	fmt.Println(i, b, f, s, i1)
}
