package main

import "fmt"

func main() {
	var n int

	var x, y, z int

	var (
		n1, n2 int
		name   string
	)

	n1, n2 = 1, 2
	name = "gopher"

	fmt.Println(n, x, y, z, n1, n2, name)
	// 0, 0, 0, 0, 1, 2, gopher
}
