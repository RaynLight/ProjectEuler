package main

import (
	"fmt"
)

func reverse(number int) int {
	reversed := 0
	for number != 0 {
		digit := number % 10
		reversed = reversed*10 + digit
		number /= 10
	}
	return reversed
}

func main() {
	var math, largestPalindrome int

	for x := 999; x > 99; x-- {
		for y := 999; y > 99; y-- {
			math = x * y
			if math == reverse(math) && math > largestPalindrome {
				largestPalindrome = math
			}
		}
	}

	fmt.Printf("Largest Palindrome: %d\n", largestPalindrome)
}
