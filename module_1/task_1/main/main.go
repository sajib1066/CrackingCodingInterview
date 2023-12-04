package main

import "fmt"

func calculateFactorial(n int) int {
	if n < 0 || n > 10 {
		return -1
	}
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

func main() {
	var n int
	fmt.Scan(&n)
	result := calculateFactorial(n)
	fmt.Println(result)
}

// time complexity: O(n)
// space complexity: O(1)
