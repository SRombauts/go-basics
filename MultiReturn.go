package main

// MultiReturn demonstrate a multi-return (public) function.
func MultiReturn(a, b int) (x, y int) {
	x = a + b
	y = a * b
	return
}


