package main

import (
	f "fmt"
)

/* Goにはコンストラクタ機能はないが、慣例的に「型のコンストラクタ」パターンを利用する */

type User struct {
	Id   int
	Name string
}

func NewUser(id int, name string) *User {
	u := new(User)
	u.Id = id
	u.Name = name
	return u
}

func main() {
	f.Println(NewUser(1, "Taro")) // コンストラクタパターン
	f.Println(&User{1, "Taro"})   // アドレス演算子
}
