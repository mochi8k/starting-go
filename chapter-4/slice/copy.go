package main

import f "fmt"

/* スライスにスライスの値を一括でコピーする. */
/* copy(dst, src []T) int */
/* copy(dst []byte, src string) int */

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 11}
	n1 := copy(s1, s2)

	f.Println("s1:", s1) // => s1: [10 11 3 4 5]
	f.Println("s2:", s2) // => s2: [10 11]
	f.Println("n1:", n1) // => n1: 2

	s3 := []int{1, 2, 3, 4, 5}
	s4 := []int{10, 11, 12, 13, 14, 15, 16}
	n2 := copy(s3, s4)

	f.Println("s3:", s3) // => s3: [10 11 12 13 14]
	f.Println("s4:", s4) // => s4: [10 11 12 13 14 15 16]
	f.Println("n2:", n2) // => n2: 5

	b := make([]byte, 9)
	n3 := copy(b, "あいうえお")
	f.Println("b:", b)   // => {バイト文字列}
	f.Println("n3:", n3) // => 9
}
