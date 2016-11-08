package main

import (
	f "fmt"
)

type Person struct {
	Id   int
	Name string
	Area string
}

func scene1() {
	/* pは*Person型 */
	p := new(Person)
	f.Println(p) // => &{0  }
}

/*
  newは基本型や参照型にも使用することができるがあまりメリットがない.
*/
func scene2() {
	i := new(int)
	f.Println(*i) // => 0

	s := new([]string)
	f.Println(*s) // => []
}

/*
  組み込み関数newを使った構造体型のポインタ生成と、
  アドレス演算子&を伴った複合リテラルによる構造体型のポインタ生成の間には、
  動作上ほとんど違いはない.
*/
func scene3() {
	type Point struct{ X, Y int }

	p1 := new(Point)
	p1.X = 1
	p1.Y = 2
	f.Println(*p1) // => {1 2}

	p2 := &Point{X: 1, Y: 2}
	f.Println(*p2) // => {1 2}
}

func main() {
	scene1()
	scene2()
	scene3()
}
