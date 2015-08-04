package main

import "fmt"

func main() {
	var c1 complex64 = 5 + 10i
	fmt.Printf("The value is: %v\n", c1)
	fmt.Printf("The real is: %g\n", real(c1))
	fmt.Printf("The imag is: %g\n", imag(c1))
}
