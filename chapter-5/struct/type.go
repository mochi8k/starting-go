package main

import (
	f "fmt"
)

/* type [定義する型] [既存の型] */
/* すでに定義されている型をもとに、新しい方を定義するための機能. */
func main() {
	scene1()
	scene2()
	scene3()
}

func scene1() {
	/* intの別名としてのMyInt型を定義 */
	type MyInt int

	var n1 MyInt = 5
	n2 := MyInt(7)
	n3 := int(8)
	f.Println(n1) // => "5"
	f.Println(n2) // => "7"
	f.Println(n3) // => "8"
}

func scene2() {
	/* example */
	type (
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float64
		IntsChannel chan []int
	)

	pair := IntPair{1, 2}
	strs := Strings{"Apple", "Banana", "Cherry"}
	amap := AreaMap{"Tokyo": {35.6, 139.6}}
	ich := make(IntsChannel)
	f.Println(pair, strs, amap, ich)
	// => [1 2] [Apple Banana Cherry] map[Tokyo:[35.6 139.6]] 0xc082012240
}

func scene3() {
	type Callback func(i int) int

	sumFn := func(ints []int, callback Callback) int {
		var sum int
		for _, v := range ints {
			sum += v
		}
		return callback(sum)
	}

	n := sumFn(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)
	f.Println(n) // => 30
}
