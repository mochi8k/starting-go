package main

import (
	"fmt"
)

func main() {
	a := 0.
	b := 72.40
	c := 072.40 // == 72.40
	d := 2.71828
	e := 1.e+0
	f := 6.67428e-11
	g := 1E6
	h := .25
	i := .12345E+5
	fmt.Println(a, b, c, d, e, f, g, h, i)

	j := 1.0e2  // == 100
	k := 1.0e+2 // == 100
	l := 1.0e-2 // == 0.01
	fmt.Println(j, k, l)

	m := 3.14
	n := int(m) // n == 3
	o := -3.14
	p := int(o) // n== -3
	fmt.Println(m, n, o, p)
}
