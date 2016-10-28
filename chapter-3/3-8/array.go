package main

import (
	"fmt"
)

func main() {

	/*
		配列型の定義
		配列型のサイズは常に固定.
	*/
	a := [5]int{}              // == [0, 0, 0, 0, 0]
	var b [5]int               // == [0, 0, 0, 0, 0]
	c := [5]int{1, 2, 3, 4, 5} // == [1, 2, 3, 4, 5]
	d := [5]int{1, 2, 3}       // == [1, 2, 3, 0, 0]
	fmt.Println(a, b, c, c[3], d)

	ia := [3]int{}        // == [0, 0, 0]
	ua := [3]uint{}       // == [0,0,0]
	ba := [3]bool{}       // == [false, false, false]
	fa := [3]float64{}    // == [0, 0, 0]
	ca := [3]complex128{} // == [(0+0i), (0+0i), (0+0i)]
	ra := [3]rune{}       // == [0, 0, 0]
	sa := [3]string{}     // == ["", "", ""]
	ni := [0]int{}        // == []
	fmt.Println(ia, ua, ba, fa, ca, ra, sa, ni)

	/*
		要素数の省略
	*/
	a1 := [...]int{1, 2, 3}       // a1 == [3]int{1, 2, 3}
	a2 := [...]int{1, 2, 3, 4, 5} // a2 == [5]int{1, 2, 3, 4, 5}
	a3 := [...]int{}              // a3 == [0]int{}
	fmt.Println(a1, a2, a3)

	/*
		要素への代入
	*/
	array := [...]int{1, 2, 3}
	array[0] = 0
	array[2] = 0
	fmt.Println(array)

	/*
		配列型の互換性
	*/
	var (
		array1 [3]int
		array2 [5]int
	)
	// array1 = array2 エラー:型が異なる代入
	fmt.Println(array1, array2)

	array3 := [3]int{1, 2, 3}
	array4 := [3]int{4, 5, 6}

	/*
		要素のコピーとなるため、array3とarray4の指すそれぞれの配列は
		メモリ上では別の領域に別れたまま.
	*/
	array3 = array4
	fmt.Println(array3) // [4, 5, 6]

	array3[0] = 0       // array4には影響しない
	array3[2] = 0       // array4には影響しない
	fmt.Println(array3) // [0, 5, 0]
	fmt.Println(array4) // [4, 5, 6]
}
