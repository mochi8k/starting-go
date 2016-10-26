package main

import (
	"fmt"
)

func main() {
	// complex128型
	c := 1.0 + 3i        // (1+3i)
	d := complex(1.0, 3) // (1+3i)
	fmt.Println(c, d)

	e := 1.3 + 4.2i // (1.3+4.2i)
	fmt.Println(e)
	fmt.Println(real(e)) // 1.3
	fmt.Println(imag(e)) // 4.2
}
