package main

type Point1 struct{ X, Y int }

func (p Point1) Set(x, y int) {
	p.X = x
	p.Y = y
}

func scene1() {
	p1 := Point1{}
	p1.Set(1, 2)
	p1.X // == 0
	p1.Y // == 0

	p2 := &Point1{}
	p2.Set(1, 2)
	p2.X // == 0
	p2.Y // == 0

}

type Point2 struct{ X, Y int }

func (p *Point2) Set(x, y int) {
	p.X = x
	p.Y = y
}

func scene2() {
	p1 := Point2{}
	p1.Set(1, 2)
	p1.X // == 1
	p1.Y // == 2

	p2 := &Point2{}
	p2.Set(1, 2)
	p2.X // == 1
	p2.Y // == 2

}

func main() {
	scene1()
	scene2()
}
