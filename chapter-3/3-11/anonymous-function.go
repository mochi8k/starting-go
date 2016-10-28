package main

import (
	"fmt"
)

func plus(x, y int) int {
	return x + y
}

var plusAlias = plus

// 関数を返す関数
func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

// 関数を引数にとる関数
func callFunction(f func()) {
	f()
}

func main() {

	/*
	   無名関数
	*/

	f1 := func(x, y int) int { return x + y }
	f1(2, 3) // == 5

	fmt.Printf("%T\n", f1) // => func(int, int) int

	var f2 func(int, int) int
	f2 = func(x, y int) int { return x + y }

	fmt.Printf("%T\n", f2) // => func(int, int) int

	fmt.Println(func(x int) int { return x }(5)) // => 5

	fmt.Println(plusAlias(10, 5)) // => 15

	rf := returnFunc()
	rf() // => "I'm a function"

	callFunction(func() {
		fmt.Println("I'm a function")
	}) // => "I'm a function"

}
