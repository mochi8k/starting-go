package main

import (
	"fmt"
)

/* fmtパッケージに定義さているStringer型 */
/*
  type Stringer interface {
	  Stringer() string
  }
*/

type T struct {
	Id   int
	Name string
}

func (t *T) String() string {
	return fmt.Sprintf("<<%d, %s>>", t.Id, t.Name)
}

func main() {
	t := &T{
		Id:   123,
		Name: "Hello",
	}
	fmt.Println(t)
}
