package main

import (
	"fmt"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	var res int64
	var max int64 = 0
	var min int64 = 1000000000
	for _, val := range arr {
		aux := int64(val)
		res += aux
		if aux > max {
			max = aux
		}
		if aux < min {
			min = aux
		}
	}
	fmt.Println(res-max, res-min)
}

// testing
func main() {
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	miniMaxSum(arr)
}
