package main

import f "fmt"

func scene1() {
	/* フィールド名のルールは、変数や関数などと同様 */
	type Person struct {
		ID   uint
		name string
		部署   string
	}

	p := Person{
		ID:   17,
		name: "山田太郎",
		部署:   "営業部",
	}

	f.Println(p)
}

func scene2() {
	/* フィールド名を省略した場合は型がフィールド名となる */
	type T struct {
		int
		float64
		string
	}

	t := T{1, 3.14, "文字列"}

	f.Println(t)
}

func scene3() {
	/* _ で無名フィールドを定義可能 */
	type T struct {
		N uint
		_ int16
		S []string
	}

	t := T{
		N: 12,
		S: []string{"A", "B", "C"},
	}

	f.Println(t) // => "{12 0 [A B C]}"
}

func main() {
	scene1()
	scene2()
	scene3()
}
