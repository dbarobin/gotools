package main

import "fmt"

func main() {

}

// 1.判断一个字符串是否为空
// if str == "" { ... }
// if len(str) == 0 {...}

// 2.判断运行 Go 程序的操作系统类型，这可以通过常量 runtime.GOOS 来判断

//if runtime.GOOS == "windows"     {
//    ...
//} else { // Unix-like
//    ...
//}

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

// 3.函数 Abs() 用于返回一个整型数字的绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 4.isGreater 用于比较两个整型数字的大小
func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}
