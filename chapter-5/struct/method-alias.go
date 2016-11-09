package main

import f "fmt"

type MyInt int

func (m MyInt) Plus(i int) int {
	return int(m) + i
}

type IntPair [2]int

func (ip IntPair) First() int {
	return ip[0]
}

type Strings []string

func (s Strings) Join(d string) string {
	sum := ""
	for _, v := range s {
		sum += v
	}
	sum += d
	return sum
}

func main() {
	f.Println(MyInt(4).Plus(2))

	ip := IntPair{1, 2}
	f.Println(ip.First())

	ss := Strings{"A", "B"}
	f.Println(ss.Join("C"))
}
