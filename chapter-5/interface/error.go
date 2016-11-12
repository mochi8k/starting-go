package main

import (
	f "fmt"
)

type error interface {
	Error() string
}

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{
		Message: "エラーが発生しました",
		ErrCode: 12345,
	}
}

func main() {
	/* errはerror型 */
	err := RaiseError()
	f.Println(err.Error()) // => "エラーが発生しました"

	/* MyError型をアサーションで取り出す */
	if e, ok := err.(*MyError); ok {
		f.Println(e.ErrCode)
	}

}
