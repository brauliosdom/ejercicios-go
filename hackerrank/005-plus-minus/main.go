package main

import (
	"fmt"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	res := []float32{0, 0, 0}
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			res[0]++
		} else if arr[i] < 0 {
			res[1]++
		} else {
			res[2]++
		}
	}
	fmt.Printf("%.6f\n", res[0]/float32(len(arr)))
	fmt.Printf("%.6f\n", res[1]/float32(len(arr)))
	fmt.Printf("%.6f\n", res[2]/float32(len(arr)))
}
