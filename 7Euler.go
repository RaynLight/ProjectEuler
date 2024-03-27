package main

import (
	"fmt"
)

func isPrime(x int) bool {
	if x <= 1 {
		return false
	}
	if x <= 3 {
		return true
	}
	if x%2 == 0 || x%3 == 0 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 || x%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	x := 1
	counter := 0
	for {
		if isPrime(x) == true {
			counter++
			if counter == 10001 {
				fmt.Println(x)
				return
			}
		}
		x += 2
	}

}
