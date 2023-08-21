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

func Pattern2(n int) {
	// given n, print the following pattern (example n=3)
	// *
	// **
	// ***
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Pattern3(n int) {
	// given n, print the following pattern (example n=3)
	// 1
	// 12
	// 123 sidenote: whenever asked to print numbers, try to avoid using 0 indexed loops.
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func Pattern4(n int) {
	// given n, print the following pattern (example n=3)
	// 1
	// 22
	// 333 sidenote: whenever asked to print numbers, try to avoid using 0 indexed loops.
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func Pattern5(n int) {
	// given n, print the following pattern (example n=3)
	// ***
	// **
	// *
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func Pattern6(n int) {
	// given n, print the following pattern (example n=3)
	// 123
	// 12
	// 1

	for i := 0; i < n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(j)
		}
		fmt.Println()

	}
}

func Pattern7(n int) {
	// given n, print the following pattern (example n=3)
	//   *
	//  ***
	// *****

	for i := 0; i < n; i++ {
		//spaces
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}

		//stars
		for j := 0; j < (2*i + 1); j++ {
			fmt.Print("*")
		}

		//spaces
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}

		fmt.Println()
	}
}
