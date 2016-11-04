package main

import (
	f "fmt"
)

func arrayPow(a [3]int) {
	/* 配列の各要素を2乗する */
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func slicePow(a []int) {
	/* 配列の各要素を2乗する */
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func main() {
	array := [3]int{1, 2, 3}
	arrayPow(array)  // 配列型の場合は値渡し
	f.Println(array) // => [1 2 3]

	slice := []int{1, 2, 3}
	slicePow(slice)  // スライスの場合は参照渡し
	f.Println(slice) // => [1 4 9]

	var (
		array1 [3]int // 宣言時にメモリが確保される
		slice1 []int  // 宣言時はnilとなる
	)

	f.Println(array1)        // => [0 0 0]
	f.Println(slice1 == nil) // => true
}
