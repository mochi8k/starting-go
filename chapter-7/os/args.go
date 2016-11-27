package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0]) // 実行したコマンド名
	fmt.Println(os.Args[1]) // 1番目のコマンドライン引数
	fmt.Println(os.Args[2]) // 2番目のコマンドライン引数
}
