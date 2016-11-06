package main

import (
	f "fmt"
)

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	/* rangeだとクローズを検出できないため注意 */
	for i := range ch {
		f.Println(i)
	}

	/*
		1
		2
		3
		deadlock!
	*/
}
