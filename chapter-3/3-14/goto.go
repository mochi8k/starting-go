package main

import (
	f "fmt"
)

func main() {

	/*
		  	ラベルとgoto文の組み合わせは、関数内に閉じたもの.
			  ブロックの内側へジャンプすることもできない.
	*/

	f.Println("A")
	goto L
	f.Println("B") // 実行されない
L: /* ラベル */
	f.Println("C")

	for {
		for {
			for {
				f.Println("start")
				goto DONE
			}
		}
	}
DONE:
	f.Println("done")

}
