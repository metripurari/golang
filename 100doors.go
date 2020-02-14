package main

import (
	"fmt"
)

func main() {
	x := [100]bool{}

	for i := 1; i <= 100; i++ {
		for j := i - 1; j < 100; j += i {
			x[j] = !x[j]
		}
	}

	for i := 0; i < 100; i++ {
		if x[i] {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
		if i%10 == 9 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}

}
