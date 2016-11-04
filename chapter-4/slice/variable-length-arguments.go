package main

/* 可変長引数(...T)は「引数の末尾に1つだけ定義できる」 */

/* ...intが可変長引数の全ての値を[]int型にまとめる. */
func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	sun(2)       // == 2
	sum(1, 2, 3) // == 6
	sum()        // == 0

	s := []int{1, 2, 3}
	sum(s...) // == 6
}
