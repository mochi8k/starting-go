package main

import (
	f "fmt"
)

func main() {
	/* panic時でもdeferは実行される */
	defer f.Println("on defer")

	panic("runtime error!") // ここでエラー終了
	f.Println("hello")
}
