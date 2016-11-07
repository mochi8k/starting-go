package main

import (
	f "fmt"
)

func scene1() {
	s := "ABC"
	&s   // 文字列型のポインタ
	s[0] // 文字列のインデックス参照(byte型)

	/*
		Goの文字列はimmutableのため、一度生成された
		文字列に対して何らかの変更を加えることは基本的に
		できないように設計されている.
	*/
	&s[0] // コンパイルエラー
}

func scene2() {
	s := ""
	for _, v := range []string{"A", "B", "C"} {
		s += v
	}

	s // == "ABC"
}

func scene3() {
	printString := func(s string) {
		f.Println(s)
	}

	s := "Hello"
	printString(s) // 文字列の実体のコピー処理は発生しない
}

func scene4() {
	/* string型の実体を表す構造体 */
	/* 型そのものにポインタを内包している */
	type stringStruct struct {
		str unsafe.Pointer // 文字列の実体へのポインタ
		len int            // 文字列のバイト長
	}
}
