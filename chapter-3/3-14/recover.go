package main

import (
	f "fmt"
)

func main() {
	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})
}

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			/* panicによるinterface{}型の値に応じて分岐 */
			switch v := x.(type) {
			case int:
				f.Printf("panic: int=%v\n", v)
			case string:
				f.Printf("panic: string=%v\n", v)
			default:
				f.Println("panic: unknown")
			}
		}
	}()
	panic(src)
}
