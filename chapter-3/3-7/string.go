package main

import (
	"fmt"
)

func main() {
	a := "Goの文字列"
	b := "\n"
	c := ""
	d := "Hello, World!\n"
	e := "日本語"
	f := "\u65e5本\U00008a9e" // 日本語
	g := "\xff\u00FF"
	h := `
    Goの
    RAW文字列リテラルによる
    複数行に渡る
    文字列
  `
	i := `
    \n
    \n
  ` // "\\n\n\\n"と同じ

	fmt.Println(a, b, c, d, e, f, g, h, i)

}
