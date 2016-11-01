package main

/* len */
/* スライスの現在の要素数を調べる */

func main() {
	s := make([]int, 8)
	len(s) // == 8

	a := [3]int{1, 2, 3}

	/* 配列にも使用可能 */
	len(a) // == 3
}
