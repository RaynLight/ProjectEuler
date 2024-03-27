package main

import "fmt"

func main() {

	var tmp int = 1
	var factor int = 1
	var number int = 600851475143

	for x := 1; x <= number; x++ {
		if number%x == 0 {
			tmp = x
			number /= x
			if tmp > factor {
				factor = tmp
			}
		}

	}
	fmt.Println(factor)
}
