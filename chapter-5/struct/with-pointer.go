package main

import (
	f "fmt"
)

type Point struct {
	X, Y int
}

func swap(p Point) {
	/* X, Yの値を入れ替える */
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func swapByPointer(p *Point) {
	/* X, Yの値を入れ替える */
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func main() {
	p := Point{X: 1, Y: 2}

	/* 構造体は値渡し */
	swap(p)
	p.X // == 1
	p.Y // == 2

	/* ポインタは参照渡し */
	swap(&p)
	p.X // == 2
	p.Y // == 1

	/* good */
	p2 := &Point{X: 1, Y: 2}
	swap(p2)
	p.X // == 2
	p.Y // == 1

}
