package main

import f "fmt"

/* select: 複数のチャネルを扱う制御構文 */

func select1() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch4 := make(chan int, 1)
	ch5 := make(chan int, 1)
	e3 := 1

	/* select文のcase節の式はすべてチャネルへの処理を伴う必要がある */
	select {

	/* ch1から受信 */
	case e1 := <-ch1:
		/* ch1からの受信が成功した場合の処理 */
		f.Println(e1)

	/* ch2から受信(2変数) */
	case e2, ok := <-ch2:
		/* ch2からの受信が成功した場合の処理 */
		f.Println(e2, ok)

	/* ch3へe3を送信 */
	case ch3 <- e3:

	/* ch5から受信したデータをch4へ送信 */
	case ch4 <- (<-ch5):

	default:
		/* case節の条件が成立しなかった場合の処理 */
	}

}

func select2() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 2

	/* 処理の継続が可能であるcase節が複数存在する場合はランダムに実行される. */
	select {
	case <-ch1:
		f.Println("ch1から受信")
	case <-ch2:
		f.Println("ch2から受信")
	case ch3 <- 3:
		f.Println("ch3へ送信")
	default:
		f.Println("ここへは到達しない")
	}

}

func select3() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	/* ch1から受信した整数を2倍してch2へ送信 */
	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)

		}
	}()

	/* ch2から受信した整数を1減算してch3へ送信 */
	go func() {
		for {
			i := <-ch2
			ch3 <- (i - 1)
		}
	}()

	n := 1
LOOP:
	for {
		select {
		/* 整数を増分させつつch1へ送信 */
		case ch1 <- n:
			n++
			/* ch3から受信した整数を出力 */
		case i := <-ch3:
			f.Println("received", i)
		default:
			if n > 20 {
				break LOOP
			}
		}
	}
}

func main() {
	select3()
}
