package main

func main() {
	var a int
	var b int32
	a = 15
	b = a + a // 编译错误
	b = b + 5 // 因为 5 是常量，所以无法通过编译
	//./type_mixing.go:7: cannot use a + a (type int) as type int32 in assignment
}
