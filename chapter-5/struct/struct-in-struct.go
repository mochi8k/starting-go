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
		Id    int
		Owner string
	}

	type A struct {
		Base /* 共通のフィールド */
		Name string
		Area string
	}

	type B struct {
		Base   /* 共通のフィールド */
		Title  string
		Bodies []string
	}

	a := A{
		Base: Base{17, "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}

	b := B{
		Base:   Base{81, "Hanako"},
		Title:  "no title",
		Bodies: []string{"A", "B"},
	}

	f.Println(a.Id)    // == 17
	f.Println(a.Owner) // == "Taro"
	f.Println(b.Id)    // == 81
	f.Println(b.Owner) // == "Hanako"
}

/* 暗黙的なフィールドの注意点 */
func scene4() {
	/*
	  ポインタ型の就職しやパッケージのプリフィックス部分は無視され、
	  純粋な型名の部分が暗黙的なフィールド名となる.
	*/
	struct {
	  T1    // フィールド名はT1
	  *T2   // フィールド名はT2
	  P.T3  // フィールド名はT3
	  *P.T4 // フィールド名はT4
	}
}

/* 再帰的な定義の禁止 */
func scene5() {
  /* 構造体のフィールドに自信の型を含むような再帰的な定義はコンパイルエラー */
  type T struct {
    T // コンパイルエラー
  }

  type T0 struct {
    T1 // T1がT0を含むのでコンパイルエラー
  }

  type T1 struct {
    T0
  }
}

func main() {
	scene1()
	scene2()
	scene3()
}
