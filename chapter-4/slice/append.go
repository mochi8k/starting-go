package main

import (
	f "fmt"
)

/* append(s S, x ...T) S */

func main() {
	append1()
	append2()
}

func append1() {
	s := []int{1, 2, 3}

	/* 4を末尾に追加 */
	s = append(s, 4)
	f.Println(s) // => [1, 2, 3, 4,]

	/* 5, 6, 7を末尾に追加 */
	s = append(s, 5, 6, 7)
	f.Println(s) // => [1, 2, 3, 4, 5, 6, 7]

	/* スライスにスライスを追加 */
	s0 := []int{1, 2, 3}
	s1 := []int{4, 5, 6}
	s2 := append(s0, s1...)
	f.Println(s2) // => [1, 2, 3, 4, 5, 6]

	/* []byte型のスライスと文字列型の場合 */
	var b []byte
	b = append(b, "あいうえお"...)
	b = append(b, "かきくけこ"...)
	b = append(b, "さしすせそ"...)
	f.Println(b) // byte値で表示される.

}

func append2() {
	/* appendとスライスの容量 */

	/* (A) */
	s := make([]int, 0, 0)
	f.Printf("(A) len=%d, cap=%d\n", len(s), cap(s))

	/* (B) */
	s = append(s, 1)
	f.Printf("(B) len=%d, cap=%d\n", len(s), cap(s))

	/* (C) */
	s = append(s, []int{2, 3, 4}...)
	f.Printf("(C) len=%d, cap=%d\n", len(s), cap(s))

	/* (D) */
	s = append(s, 5)
	f.Printf("(D) len=%d, cap=%d\n", len(s), cap(s))

	/* (E) */
	s = append(s, 6, 7, 8)
	f.Printf("(E) len=%d, cap=%d\n", len(s), cap(s))

	/* (F) */
	s = append(s, 9)
	f.Printf("(F) len=%d, cap=%d\n", len(s), cap(s))

	/*
		  容量が不足したタイミングで、Goのランタイムが、スライスの容量を増やす.
			(A) len=0, cap=0
			(B) len=1, cap=1
			(C) len=4, cap=4
			(D) len=5, cap=8
			(E) len=8, cap=8
			(F) len=9, cap=16
	*/

	/*
	   	容量が不足した場合に、どの程度の容量が割り当てられるかは、
	  	Goのランタイムに依存する.
	*/

	s1 := make([]int, 1024, 1024)
	f.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
	// => len=1024, cap=1024

	s1 = append(s1, 0)
	f.Printf("len=%d, cap=%d\n", len(s1), cap(s1))
	// => len=1025, cap=1312

}
