package main

/*
  同じパッケージであれば、ファイルが分割されていても、
  定数や変数といった要素については相互に参照可能.
  しかし、インポート宣言のみ各ファイルで独立している.
*/

import (
	"fmt"
)

func main() {
	fmt.Println("AppVersion:", AppVersion)
	printMessage("Hello!")
}
