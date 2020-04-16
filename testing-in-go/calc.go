package main

import (
	"fmt"
)

func Addition(x, y int) (result int) {
	return x + y
}

func main() {
	fmt.Printf("The sum of %d and %d is %d", 5, 3, Addition(5, 3))
}
