package main

import (
	f "fmt"
)

/* len: チャネルのバッファ内に溜まっているデータ数. */

func main() {
	ch := make(chan string, 3)
	ch <- "Apple"
	len(ch) // == 1
	ch <- "Banana"
	ch <- "Cherry"
	len(ch) // == 3

	/* Bad */
	if len(ch) > o {
		i := <-ch // len(ch) > 0は保証さない.
	}
}
