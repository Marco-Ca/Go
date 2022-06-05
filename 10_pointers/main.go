package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println("a: ", a, "b: ", b)
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println("b: ", *b)
}
