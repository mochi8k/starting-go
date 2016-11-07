package main

func main() {
	type T0 int
	type T1 int

	t0 := T0(5)   // t0 == 5
	i0 := int(t0) // i0 == 5

	t1 := T1(8)   // t0 == 8
	i1 := int(t1) // i0 == 8

	/* 同じ型から派生した場合でも、エイリアス間には互換性が成り立たない. */
	t0 = t1 // コンパイルエラー
}
