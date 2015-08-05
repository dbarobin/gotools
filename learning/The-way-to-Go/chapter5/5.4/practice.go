package main

import "fmt"

func main() {
	func1()
	//func2()
	//func3()
	func4()
	func5()
}

func func1() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
}

func func2() {
	for i := 0; ; i++ {
		fmt.Println("Value of i is now:", i)
	}
}

func func3() {
	for i := 0; i < 3; {
		fmt.Println("Value of i:", i)
	}
}

func func4() {
	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
}

func func5() {
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
