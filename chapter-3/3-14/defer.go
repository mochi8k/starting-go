package main

import (
	f "fmt"
	"os"
)

func runDefer() {
	/*
	  deferに登録された式は関数の終了時に評価される.
	  登録された式は「あとで登録された式」から評価される.
	  登録できる式は関数呼び出しのみ.
	*/
	defer f.Println("1")
	defer f.Println("2")
	defer f.Println("3")
	f.Println("done")
}

func main() {
	runDefer()
	// done
	// 3
	// 2
	// 1

	defer func() {
		f.Println("A")
		f.Println("B")
		f.Println("C")
	}() // 関数呼び出しの形式

	file, err := os.Open("/path/to/file")
	if err != nil {
		/* エラー処理 */
		return
	}

	/* ファイルのクローズ処理を登録 */
	defer file.Close()

}
