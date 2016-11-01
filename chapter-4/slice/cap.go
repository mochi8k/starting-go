package main

import (
	f "fmt"
)

/* cap */
/* スライスの現在の容量数を調べる */

func main() {

	/* 要素数5, 容量5のスライス */
	s1 := make([]int, 5)
	f.Println(len(s1)) // =>  5
	f.Println(cap(s1)) // =>  5

	/* 要素数5, 容量10のスライス */
	s2 := make([]int, 5, 10)
	f.Println(len(s2)) // =>  5
	f.Println(cap(s2)) // =>  10

}
