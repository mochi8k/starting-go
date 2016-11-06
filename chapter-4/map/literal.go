package main

import (
	f "fmt"
)

func main() {
	m1 := map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	f.Println(m1) // map[1:Taro 2:Hanako 3:Jiro]

	m2 := map[int]string{
		1: "Taro",
		2: "Hanako",
		3: "Jiro", // カンマが必要
	}
	f.Println(m2) // map[1:Taro 2:Hanako 3:Jiro]

	m3 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}
	f.Println(m3) // map[1:[1] 2:[1 2] 3:[1 2 3]]

	m4 := map[int][]int{
		/* 内側のスライスリテラルの[]intは省略可能 */
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	f.Println(m4) // map[1:[1] 2:[1 2] 3:[1 2 3]]

	m5 := map[int]map[float64]string{
		/* 内側のマップリテラルは省略可能 */
		1: {3.14: "円周率"},
	}
	f.Println(m5) // map[1:map[3.14:円周率]]
}
