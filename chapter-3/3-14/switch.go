package main

import (
	f "fmt"
)

func main() {

	/* 式によるswitch */
	n := 3
	switch n {
	case 1, 2:
		f.Println("1 or 2")
		/* 次のcase節へフォールスルーしない */
	case 3, 4:
		f.Println("3 or 4")

		/* フォールスルーは明示的に記述する. */
		fallthrough
	case 5, 6:
		f.Println("5 or 6")
	default:
		f.Println("unknown")
	}

	/* 型 */
	switch n := 3; n {
	case 1:
		f.Println("one")
	case 2 + 0i:
		f.Println("two")
	case 3.0:
		f.Println("three")
	case "4": /* 互換性のない型を定義するとエラー */
		f.Println("four")
	case 4.5: /* 整数と互換がないためエラー */
		f.Println("4.5")
	}

	/* 式を伴うcase */
	n := 4
	switch { /* switch true と同じ意味となる */
	case n > 0 && n < 3:
		f.Println("0 < n < 3")
	case n > 3 && n < 6:
		f.Println("3 < n < 6")
	}

	/* 型アサーション(動的に変数の型をチェックする) interface型の変数.(型) */
	var x1 interface{} = 3
	i1 := x1.(int)     // 変数iはint型で値は3
	f1 := x1.(float64) // エラーが発生

	/* 第二引数までを代入した場合はエラーが発生しない. */
	var x2 interface{} = 3.14
	i2, isInt := x2.(int) // i2 == 0, isInt == false
	f2, isFloat64 := x2.(float64) // f2 == 3.14, isFloat64 == true
	s, isString := x2.(string) // s == "", isString == false

	var x3 interface{} = "hello"
	if x3 == nil {
		// dosomething
	} else if i, isInt := x.(int); isInt {
		// dosomething
	} else if s, isString := x.(string); isString {
		// dosomething
	} else {
		// dosomething
	}
}
