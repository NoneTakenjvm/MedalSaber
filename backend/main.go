package main

import "fmt"

func main() {

	x := 10
	byPointer(&x)
	fmt.Println(x)
	// would equal 20

	x = 10
	x = byValue(x)
	fmt.Println(x)
	// would also equal 20
}

func byPointer(x *int) {
	*x *= 2
}

func byValue(x int) int {
	return x * 2
}
