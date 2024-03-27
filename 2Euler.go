package main

import "fmt"

func main() {
	sum := 2 // start with 2 because you've already initialized sum to 2
	num1 := 1
	num2 := 2
	num3 := 0

	for {
		num3 = num1 + num2
		if num3 >= 4000000 {
			break
		}
		if num3%2 == 0 {
			sum += num3
		}
		num1 = num2
		num2 = num3
	}
	fmt.Println(sum)
}
