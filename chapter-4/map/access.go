package main

import (
	f "fmt"
)

func main() {
	m1 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	s1 := m[1] // s == "A"
	s1 := m[9] // s == "" 型の初期値が設定される.

	m2 := map[int]int{1: 0}
	i1 := m2[7] // i == 0 定義されているのか初期値なのかわからない.

	m3 := map[int]string{
		1: "A",
		2: "B",
		3: "",
	}

	/* 第二引数に存在するかどうかが返却される. */
	/* 第二引数の変数名を ok とするのはGoにおける一種のイディオム */
	s3, ok := m3[1] // s == "A", ok == true
	s3, ok := m3[9] // s == "", ok == false
	s3, ok := m3[3] // s == "", ok == true

	m4 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	if _, ok := m4[1]; ok {
		/* dosomething */
	}

}
