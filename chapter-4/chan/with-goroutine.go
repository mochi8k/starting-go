package main

import (
	f "fmt"
)

func receiver(ch <-chan int) {
	for {
		i := <-ch
		f.Println(i)
	}
}

func main() {
	ch := make(chan int)

	go receiver(ch)

	i := 0
	for i < 20 {
		ch <- i
		i++
	}

}

func main2() {

	ch2 := make(chan rune, 3)

	ch2 <- 'A'
	ch2 <- 'B'
	ch2 <- 'C'
	// ch2 <- 'D' デッドロック発生

}
