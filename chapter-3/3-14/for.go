package main

import (
	f "fmt"
)

func main() {

	/* 無限ループ */
	for {
		f.Println("I'm in infinite loop")
	}

	/* break */
	i1 := 0
	for {

		f.Println(i1)
		i1++

		if i1 == 100 {
			break
		}

	}

	/* 条件付きfor */
	i2 := 0
	for i < 100 {
		f.Println(i2)
		i++
	}

	/* 古典的for */
	for i := 0; i < 100; i++ {
		f.Println(i)
		i++
	}

	/* continue */
	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			continue
		}
		f.Println(i)
		i++
	}

	/* 範囲節によるfor */
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i, v := range fruits {

		// i: インデックス
		// v: 値

		fmt.Println(i, v)
	}
}
