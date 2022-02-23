package main

import "fmt"

func main() {
	fmt.Println("Fibonacci of 6: ", fibonacci(6))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
