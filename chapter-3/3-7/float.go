package main

import (
	"fmt"
)

func main() {
	i := 1              // int
	f64 := 1.0          // float64
	f32 := float32(1.0) // float32

	fmt.Println(i, f64, f32)

	zero := 0.0
	pinf := 1.0 / zero  // +Inf 正の無限大
	ninf := -1.0 / zero // -Inf 負の無限大
	nan := zero / zero  // NaN 非数

	fmt.Println(zero, pinf, ninf, nan)
}
