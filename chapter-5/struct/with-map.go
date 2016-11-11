package main

type User struct {
	Id   int
	Name string
}

func main() {
	/* キーが構造体のマップ */
	m1 := map[User]string{
		{Id: 1, Name: "Taro"}: "Tokyo",
		{Id: 2, Name: "Jiro"}: "Osaka",
	}

	/* キーが構造体のマップ */
	m2 := map[User]string{
		1: {Id: 1, Name: "Taro"},
		2: {Id: 2, Name: "Jiro"},
	}

	/* 値がスライスのマップ */
	ms := map[int][]string{
		1: {"A", "B", "C"},
	}

	/* 値がマップのマップ */
	mm := map[int]map[int]string{
		1: {1: "Apple", 2: "Banana", 3: "Cherry"},
	}

}
