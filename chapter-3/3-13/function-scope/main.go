package main

import (
	f "fmt"
)

/*
  関数の中で定義された変数や定数は、
  定義された関数の中でのみ参照可能.
*/

func appName() string {
	const AppName = "Go Application"
	var Version = "1.0"
	return AppName + " " + Version
}

func doSomething(a int) (b string) {

	b = "hello1"
	// var a int          // 定義済みのためエラー
	// const b = "string" // 定義済みのためエラー

	{
		/* 関数より深いブロック */
		var a int
		b = "hello2" // 参照可能

		const b = "hello3" // 宣言可能
		f.Println(a)

		// return // エラー
		// return b // b == "hello3"
	}

	return b // b == "hello2"
}

func main() {
	f.Println(appName())
	// f.Println(AppName) // コンパイルエラー
	f.Println(doSomething(5)) // => hello2
}
