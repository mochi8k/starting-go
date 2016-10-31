package main

import (
	f "fmt"
)

/*
  go文は並列を司る機能.
  defer文と同様に、関数呼び出し形式の式を受け取る.
*/

func sub() {
	for {
		f.Println("second loop")
	}
}

func main() {
	go sub()

	go func() {
		for {
			f.Println("third loop")
		}
	}()

	for {
		f.Println("first loop")
	}
}
