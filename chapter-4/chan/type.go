package main

/* chan[データ型] */

func main() {
	/* chan   双方向チャネル */
	var ch0 chan int

	/* <-chan 受信専用チャネル */
	var ch1 <-chan int

	/* chan<- 送信専用チャネル */
	var ch2 chan<- int

	/* 代入の可否 */
	ch1 = ch0 /* OK */
	ch2 = ch0 /* OK */
	ch0 = ch1 /* NG */
	ch0 = ch2 /* NG */
	ch1 = ch2 /* NG */
	ch2 = ch1 /* NG */

}
