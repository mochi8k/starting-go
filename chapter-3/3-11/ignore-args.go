package main

type T struct {
	Value int
}

type I interface {
	// 引数が2つ必要であると定義
	RequiredFunction(a, b int) int
}

func (*T) RequiredFunction(a, _ int) int {
	// 実装に第2引数は不要なためブランクにする.
	return a
}
