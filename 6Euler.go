package main

import (
	"fmt"
)

func squares() int {
	nums := 0
	for x := 1; x < 101; x++ {
		nums += x * x
	}

	return nums
}

func add_thenSquare() int {
	nums := 0
	for x := 1; x < 101; x++ {
		nums += x
	}
	return nums * nums
}

func main() {
	var all_squared int = squares()
	var added_then_squared int = add_thenSquare()
	fmt.Println(added_then_squared - all_squared)
}
