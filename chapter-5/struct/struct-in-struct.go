package main

import f "fmt"

func scene1() {
	type Feed struct {
		Name   string
		Amount uint
	}

	type Animal struct {
		Name string
		Feed Feed /* Feed型の埋め込み */
	}

	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	f.Println(a) // => "{Monkey {Banana 10}}"
}

func scene2() {
	type Feed struct {
		Name   string
		Amount uint
	}

	type Animal struct {
		Name string
		Feed /* フィールド名を省略した構造体の埋め込み */
	}

	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	/*
	  フィールド名を省略して埋め込まれた構造体のフィールド名が「一意に」定まる限り、
	  中間のフィールド名を省略してアクセス可能.
	*/
	f.Println(a.Amount) // => 10
	a.Amount = 15
	f.Println(a.Amount)      // => 15
	f.Println(a.Feed.Amount) // => 15
}

func scene3() {

  /* 「異なる構造体型に共通の性質を持たせる」構造体 */
  type Base struct {
    Id int
    Owner string
  }

  type A struct {
    Base /* 共通のフィールド */
    Name string
    Area string
  }

  type B struct {
    Base /* 共通のフィールド */
    Title string
    Bodies []string
  }

  a := A{
    Base: Base{17, "Taro"},
    Name: "Taro",
    Area: "Tokyo",
  }

  b := B{
    Base: Base{81, "Hanako"},
    Title: "no title",
    Bodies: []string{"A", "B"},
  }

  f.Println(a.Id) // == 17
  f.Println(a.Owner) // == "Taro"
  f.Println(b.Id) // == 81
  f.Println(b.Owner) // == "Hanako"
}

func main() {
	scene1()
	scene2()
  scene3()
}
