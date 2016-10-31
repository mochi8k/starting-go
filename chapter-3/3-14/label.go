package main

import (
	f "fmt"
)

func main() {

LOOP:
	for {
		for {
			for {
				f.Println("start")
				break LOOP
			}
			f.Println("ここは通らない")
		}
		f.Println("ここは通らない")
	}
	f.Println("finish")

L:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j > 1 {
				continue L
			}
			f.Println(i, j)
		}
		f.Println("ここは通らない")
	}

}
