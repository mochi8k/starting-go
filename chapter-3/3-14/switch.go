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

}
