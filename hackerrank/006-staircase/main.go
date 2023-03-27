package main

import (
	"fmt"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	for i := 0; i < int(n); i++ {
		for j := int(n - 1); j > i; j-- {
			fmt.Print(" ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}
}
