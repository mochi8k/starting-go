package main

import (
	f "fmt"
)

func main() {
	s := []string{"Apple", "Banana", "Cherry"}

	for i, v := range s {
		f.Printf("[%d] => %s\n", i, v)
		s = append(s, "Melon") // 配列の要素を追加してもループは3回
	}

	for i := 0; i < len(s); i++ {
		f.Printf("[%d] => %s\n", i, s[i])
		s = append(s, "Melon") // 無限ループとなる
	}
}
