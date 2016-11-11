package main

import (
	f "fmt"
)

type Point struct{ X, Y int }

func scene1() {
	ps := make([]Point, 5)
	for _, p := range ps {
		f.Println(p.X, p.Y)
	}
}

type Points []*Point

func (ps Points) ToString() string {
	str := ""

	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += f.Sprintf("[%d,%d]", p.X, p.Y)
		}
	}
	return str
}

func scene2() {
	ps := Points{}
	ps = append(ps, &Point{X: 1, Y: 2})
	ps = append(ps, nil)
	ps = append(ps, &Point{X: 3, Y: 4})
	f.Println(ps.ToString())
}

func main() {
	scene1()
	scene2()
}
