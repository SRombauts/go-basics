// Package multi demonstrates how to organize sources as packages
package multi

// Return demonstrates a multi-return (public) function.
func Return(a, b int) (x, y int) {
	x = a + b
	y = a * b
	return
}
