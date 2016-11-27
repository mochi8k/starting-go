package main

/*
  "os"はGoが動作する各OSに依存した機能群を内部に隠蔽した、
  プラットフォームに独立したAPIを提供する.
*/
import (
	"fmt"
	"log"
	"os"
)

/* os.Exit */
func exit() {

	/* 破棄される */
	defer func() {
		fmt.Println("defer")
	}()

	/* プログラムを停止する(deferは実行されない). */
	os.Exit(0)
}

/* log.Fatal */
func fatal() {

	/* 破棄される */
	defer func() {
		fmt.Println("defer")
	}()

	_, err := os.Open("foo")
	if err != nil {
		/* 引数として与えた値を標準出力へ出力し、os.Exit(1)を実行する. */
		log.Fatal(err)
	}
}

func main() {
	exit()
	fatal()
}
