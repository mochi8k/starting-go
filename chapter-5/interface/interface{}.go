package main

import (
	"fmt"
)

type T struct {
	id int
}

func (t *T) GetId() int { return t.id }

/*
interfaceもtypeと同様に、エイリアスを定義せずにむき出しの型定義として使用できる.
interface{}型というのは、空のインターフェースを表すため、どんな型でも受け取ることができる.
*/

func ShowId(id interface {
	GetId() int
}) {
	/* 変数idは「GetId() int」というメソッドを持つ型 */
	fmt.Println(id.GetId())
}

func main() {
	t := &T{id: 1234}
	ShowId(t)
}
