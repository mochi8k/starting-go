package main

import (
	f "fmt"
)

/*
  パッケージの初期化用関数.
  main関数よりも先に実行される.
  init関数は複数定義することが可能であり、定義順に実行される.
*/
func init() {
	f.Println("init1()")
}

func init() {
	f.Println("init2()")
}

func init() {
	f.Println("init3()")
}

func main() {
	f.Println("main()")
}
