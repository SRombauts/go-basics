// Simple tests of the Go 1.1 language.
package main

import "fmt"

// MultiReturn demonstrate a multi-return (public) function.
func MultiReturn(a, b int) (x, y int) {
	x = a + b
	y = a * b
	return
}

// main Entry point of the application.
func main() {
	fmt.Println("Hello World!")
	x, y := MultiReturn(7, 6)
	fmt.Println("MultiReturn:", x, y)
}
