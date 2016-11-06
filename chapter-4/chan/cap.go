package main

/* cap: チャネルのバッファサイズを取得 */

func main() {
	ch1 := make(chan string)
	cap(ch1) // == 0

	ch2 := make(chan string, 3)
	cap(ch2) // == 3
}
