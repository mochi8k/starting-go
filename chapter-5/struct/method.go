package main

import f "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Render() {
	f.Printf("<%d,%d>", p.X, p.Y)
}

func (p *Point) Sum() int {
	return p.X + p.Y
}

func scene1() {

}

func main() {
	p := &Point{X: 5, Y: 12}
	p.Render()         // <5,12>
	f.Println(p.Sum()) // 17
}
