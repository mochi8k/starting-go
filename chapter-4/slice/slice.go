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

	/* スライスではなく配列 */
	var a1 [5]int
	f.Println(a1) // [0 0 0 0 0]

}
