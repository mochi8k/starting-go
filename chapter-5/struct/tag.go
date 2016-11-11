package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func scene1() {
	/* タグ(Tag): フィールドにメタ情報を付与する機能 */
	type User struct {
		Id   int    "ID"
		Name string "名前"
		Age  uint   "年齢"
	}

	u := User{Id: 1, Name: "Taro", Age: 32}

	/* reflect.Type型 */
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
	}
}

func scene2() {
	type User struct {
		Id   int    `json:"user_id"`
		Name string `json:"user_name"`
		Age  uint   `json:"user_age"`
	}

	u := User{Id: 1, Name: "Taro", Age: 32}
	bs, _ := json.Marshal(u)
	fmt.Println(string(bs))
}

func main() {
	scene1()
	scene2()
}
