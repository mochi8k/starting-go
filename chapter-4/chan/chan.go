package main

import (
	f "fmt"
)

/*
チャネルは「キュー(FIFO)」の性質を備えるデータ構造.
バッファとはこのキューを格納する領域であり、
バッファサイズとはこのキューのサイズである.
*/

func main() {
	/* バッファサイズ0のチャネル */
	ch1 := make(chan int)
	f.Println(ch1) // 0xc42006a060

	/* バッファサイズのチャネル */
	ch2 := make(chan int, 8)
	f.Println(ch2) // 0xc420076000

	ch3 := make(chan int, 10)

	/* 整数5を送信 */
	ch3 <- 5

	/* 整数値を受信 */
	i := <-ch3
	f.Println(i) // 5
}
