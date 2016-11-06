package main

import (
	f "fmt"
)

/* map[キーの型]要素の型 */

func main() {
	var m1 map[int]string
	m1 = make(map[int]string)

	m1[1] = "US"
	m1[81] = "Japan"
	m1[86] = "China"

	f.Println(m1) // map[1:US 81:Japan 86:China]

	m2 := make(map[string]string)

	m2["Yamada"] = "Taro"
	m2["Sato"] = "Hanako"
	m2["Yamada"] = "Jiro"

	f.Println(m2) // map[Yamada:Jiro Sato:Hanako]

	m3 := make(map[float64]int)

	/* 丸められて全て1.0となる.*/
	m3[1.0000000000000000000001] = 1
	m3[1.0000000000000000000002] = 2
	m3[1.0000000000000000000003] = 3

	f.Println(m3) // map[1:3]
}
