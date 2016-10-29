package foo

/* 定数 */
const (
	MAX            = 100 // 公開される
	internal_const = 1   // パッケージ内のみ
)

/* パッケージ変数 */
var (
	N = 512 // 公開される
	m = 256 // パッケージ内のみ
)

/* 公開される */
func FooFunc(n int) int {
	return internalFunc(n)
}

/* パッケージ内のみ */
func internalFunc(n int) int {
	return n + 1
}
