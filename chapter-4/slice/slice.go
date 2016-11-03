package main

import (
	f "fmt"
)

/* スライス: 可変長配列を表現する型 */

func main() {

	/* int型のスライス */
	var s1 []int

	f.Println(s1) // => []

	/* makeによる生成 */
	s2 := make([]int, 5)

	f.Println(s2) // [0 0 0 0 0]

	/* リテラルでの生成 */
	s3 := []int{1, 2, 3, 4, 5}
	len(s3) // == 5
	cap(s3) // == 5

	/* スライスではなく配列 */
	var a1 [5]int
	f.Println(a1) // [0 0 0 0 0]

	/* 要素への代入と参照 */
	fs := make([]float64, 3) // == [0, 0, 0]
	fs[0] = 3.14
	fs[1] = 6.28     // == [3.14, 6.28, 0]
	f.Println(fs[4]) // ランタイムパニック
}
