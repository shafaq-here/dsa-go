package main

import "fmt"

func Pattern1(n int) {
	// given n, print the following pattern (example n=3)
	// ***
	// ***
	// ***

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
