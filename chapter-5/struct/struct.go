package main

type Point struct {
	X int
	Y int
}

func main() {
	scene1()
	scene2()
}

func scene1() {
	var pt Point
	pt.X // == 0
	pt.Y // == 0

	pt.X = 10
	pt.Y = 8

	pt.X // == 10
	pt.Y // == 8
}

func scene2() {
	/* 構造体のフィールド定義の順序を完全に一致させる必要がある */
	pt1 := Point{1, 2}
	pt1.X // == 1
	pt1.Y // == 2

	/* フィールドを明示的に指定する */
	pt2 := Point{X: 1, Y: 2}
	pt2.X // == 1
	pt2.Y // == 2
}
