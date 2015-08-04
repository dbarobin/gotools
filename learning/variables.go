package main

import "fmt"
import "os"

func main() {
	//var identifier [type] = value
	var a int = 15
	var i = 5
	var b bool = false
	var str string = "Go says hello to the world!"

	var aa = 15
	var bb = false
	var strstr = "Go says hello to the world!"

	var (
		aaa       = 15
		bbb       = false
		strstrstr = "Go says hello to the world!"
		numShips  = 50
		city      string
	)

	var n int64 = 2

	var (
		HOME   = os.Getenv("Home")
		USER   = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)
}
