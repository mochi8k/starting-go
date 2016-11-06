package main

import (
	f "fmt"
)

func main() {
	m := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	delete(m, 2)

	f.Println(m) // map[1:A 3:C]
}
