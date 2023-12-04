package main

import "fmt"

func calculateNaturalNumber(a int, b int) int {
	var result, temp int
	if a > b {
		temp = a
		a = b
		b = temp
	}
	// for i := a; i <= b; i++ {
	// 	result += i
	// }
	result = ((b - a + 1) * (a + b)) / 2
	return result
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result := calculateNaturalNumber(a, b)
	fmt.Println(result)
}

// time complexity: O(1)
// space complexity: O(1)
