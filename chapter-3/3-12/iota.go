package main

const (
	A = iota // == 0
	B = iota // == 1
	C = iota // == 2
)

const (
	D = iota // == 0
	E        // == 1
	F        // == 2
)

const (
	G = 1 + iota // == 1
	H            // == 2
	I            // == 3
	J = iota     // == 3
)

/*
  iotaは参照の有無によらずconstブロックの中で
  定数が定義されるたびに1ずつ増分する.
  カウンターのようなイメージではなく、インデックスのイメージ.
*/
const (
	K = iota // == 0
	L        // == 1
	M        // == 2
	N = 17   // == 17
	O = iota // 4
	P        // 5
)
