package main

import "fmt"

func main() {
	a := 5.0
	b := int(a)
	fmt.Println(b)

	var aa = 5
	c := int(aa)
	d := int(c)

	fmt.Println(c)
	fmt.Println(d)
}
