package main

import (
	f "fmt"
)

func scene1() {
	p := &[3]int{1, 2, 3}
	f.Println((*p)[0]) // => "1"
}

func scene2() {
	a := [3]string{"Apple", "Banana", "Cherry"}
	p := &a

	/* コンパイラが (*a)[1] に変換してくれる */
	f.Println(a[1]) // => "Banana"
	f.Println(p[1]) // => "Banana"

	p[2] = "Grape"
	f.Println(a[2]) // => "Grape"
	f.Println(p[2]) // => "Grape"

}

func scene3() {
	p := &[3]int{1, 2, 3}

	/* デリファレンスを省略できる */
	f.Println(len(p)) // => "3"
	f.Println(cap(p)) // => "3"
	f.Println(p[0:2]) // => "[1 2]"
}

func scene4() {
	p := [3]string{"Apple", "Banana", "Cherry"}
	for i, v := range p {
		f.Println(i, v)
	}
}

func main() {
	scene1()
	scene2()
	scene3()
	scene4()
}
