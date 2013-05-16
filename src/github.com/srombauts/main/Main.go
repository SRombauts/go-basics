// Simple tests of the Go 1.1 language.
package main

import "fmt"

// main Entry point of the application.
func main() {
	fmt.Println("Hello World!")
	x, y := MultiReturn(7, 6)
	fmt.Println("MultiReturn:", x, y)
}
