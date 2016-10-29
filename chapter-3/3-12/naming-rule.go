package main

import (
	"fmt"
)

/*
  Goの識別子に利用できる文字種は、
  「Unicodeにおいて『文字』と定義されたもの」 &&
  「Unicodeにおいて『数字』と定義されたもの」 &&
  「_」アンダースコア
*/

/*
  日本語を使った例
*/
const (
	朝の挨拶 = "おはよう"
	昼の挨拶 = "こんにちは"
	夜の挨拶 = "こんばんは"
)

func あいさつ(m string) {
	fmt.Println(m)
}

func main() {
	あいさつ(昼の挨拶) // => "こんにちは"
}
