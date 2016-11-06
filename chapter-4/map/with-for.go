package main

import (
	f "fmt"
)

func main() {
	m1 := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}

	/* キーの順序は保証されない */
	for k, v := range m1 {
		f.Printf("%d => %s\n", k, v)
	}

	/*
		1 => Apple
		2 => Banana
		3 => Cherry
	*/

	m2 := map[string]int{
		"Apple":  88,
		"Banana": 107,
		"Cherry": 46,
	}

	m2["Grape"] = 66
	m2["Lemon"] = 16
	m2["Orange"] = 44
	m2["Pineallple"] = 73

	for k, v := range m2 {
		f.Printf("%s => %d\n", k, v)
	}

	/*
	   Pineallple => 73
	   Apple => 88
	   Banana => 107
	   Cherry => 46
	   Grape => 66
	   Lemon => 16
	   Orange => 44
	*/
}
