package main

func main() {
	m := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	len(m) // == 3

	m[4] = "D"

	len(m) // == 4
}
