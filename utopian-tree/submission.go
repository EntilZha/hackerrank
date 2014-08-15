package main

import (
	"fmt"
)

func main() {
	var nTests int
	testCases := make([]int, 0)
	fmt.Scanf("%d", &nTests)
	for i := 0; i < nTests; i++ {
		var testCase int
		fmt.Scanf("%d", &testCase)
		testCases = append(testCases, testCase)
	}
	for _, n := range testCases {
		fmt.Println(calculateHeight(n))
	}
}

func calculateHeight(n int) int {
	return calculateHeightRecursive(0, n, 1)
}
func calculateHeightRecursive(cycle int, n int, height int) int {
	if cycle >= n {
		return height
	}
	if cycle%2 == 0 {
		return calculateHeightRecursive(cycle+1, n, 2*height)
	} else {
		return calculateHeightRecursive(cycle+1, n, height+1)
	}
}
