package main

import (
	"fmt"
)

func main() {
	var (
		// 符号あり
		i8  int8  = 127
		i16 int16 = 32767
		i32 int32 = 2147483647
		i64 int64 = 9223372036854775807

		// 符号なし
		u8  uint8  = 255
		u16 uint16 = 65535
		u32 uint32 = 4294967295
		u64 uint64 = 18446744073709551615

		// 実装依存
		i  int  = 1
		ui uint = 1

		// ポインタの値をそのまま格納するのに十分な大きさの符号なし整数
		uip uintptr = 18446744073709551615
	)

	fmt.Println(i8, i16, i32, i64)
	fmt.Println(u8, u16, u32, u64)
	fmt.Println(i, ui)
	fmt.Println(uip)

}
