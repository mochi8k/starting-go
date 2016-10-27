package main

import (
	"fmt"
)

func main() {

	a := [5]int{} // 0で初期化される.
	var b [5]int
	c := [5]int{1, 2, 3, 4, 5}
	d := [5]int{1, 2, 3} // 足りない分は0で初期化される.
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
}
