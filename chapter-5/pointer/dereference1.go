package main

import (
	f "fmt"
)

func inc(p *int) {
	/* pをデリファレンスして+1増分 */
	*p++
}

func main() {
	i := 1

	/*
	   アドレス演算子を使うことで、プリミティブ型の値を
	   参照渡しで関数を呼び出す.
	*/
	inc(&i)
	inc(&i)
	inc(&i)

	f.Println(i) // => "4"
}
