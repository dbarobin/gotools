package main

import "fmt"
import "time"

func ready(w string, sec int64) {
	secs := time.Duration(sec) * time.Second
	time.Sleep(secs)
	fmt.Println(w, "is ready!")
}

func main() {
	var sec int
	go ready("Tee", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	secs := time.Duration(sec) * time.Second
	time.Sleep(5 * secs)
}
