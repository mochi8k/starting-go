package main

import (
	f "fmt"
)

func showStruct(s struct{ X, Y int }) {
	f.Println(s)
}

func main() {
	s := struct{ X, Y int }{X: 1, Y: 2}
	showStruct(s) // "{1 2}"
}
