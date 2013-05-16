// Simple tests of the Go 1.1 language.
package main

import (
	"fmt"
	"github.com/srombauts/multi"
)

// main Entry point of the application.
func main() {
	fmt.Println("Hello World!")
	x, y := multi.Return(7, 6)
	fmt.Println("multi.Return:", x, y)
}
