package main

import (
	e "errors"
	f "fmt"
)

func main() {

	x := 1

	/* 基本形 */
	if x == 1 {
		f.Println("true")
	}

	/* if else */
	if x == 1 {
		// dosomething
	} else if x > 10 {
		// dosomething
	} else {
		// dosomething
	}

	/* ifの条件式にはbool型のみ記述可能 */
	if true {
		// dosomething
	}

	/* 条件式がbool型ではないためエラーとなる */
	if 1 {
		// dosomething
	}

	/*
	  	簡易文(Simple Statement)付きif
	    ここで宣言した変数はifブロック内のスコープとなる.
	*/
	if x, y := 3, 5; x < y {
		// dosomething
	}

	if _, err := errorGenerator(); err != nil {
		// dosomething
	}

}

func errorGenerator() (string, error) {
	return "", e.New("Error")
}
