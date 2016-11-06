package main

/* close([チャネル]) */
/* 指定したチャネルをクローズする. */

func main() {

	/* クローズ状態のチャネルに対する送信 */
	ch1 := make(chan int, 1)
	close(ch1)
	ch1 <- 1 // ランタイムパニック

	/* クローズ状態のチャネルに対する受信 */
	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch2)

	// ランタイムパニックにはならない
	<-ch2 // == 1
	<-ch2 // == 2
	<-ch2 // == 3
	<-ch2 // == 0 // 型の初期値
	<-ch2 // == 0 // 型の初期値

	/* クローズ状態の検出 */
	ch3 := make(chan int, 3)
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3

	close(ch3)

	// バッファ内が空, かつクローズされた状態の場合: false */
	i, ok := <-ch3 // i == 1, ok == true
	i, ok := <-ch3 // i == 2, ok == true
	i, ok := <-ch3 // i == 3, ok == true
	i, ok := <-ch3 // i == 0, ok == false
}
