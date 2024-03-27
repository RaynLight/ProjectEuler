package main

import (
	"fmt"
)

func mod20(x int) bool {
	value := true
	for num := 20; num > 0; num-- {
		if x%num != 0 {
			value = false
		}
	}
	return value
}

func main() {
	var x int = 2
	for {
		if mod20(x) == true {
			fmt.Println(x)
			return
		}
		x++
	}

}
