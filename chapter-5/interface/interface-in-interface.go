package main

type I0 interface {
	Method1() int
}

type I1 interface {
	I0
	Method2() int
}

type I2 interface {
	I1
	Method3() int
}
