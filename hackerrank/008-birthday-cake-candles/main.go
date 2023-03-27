package main

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var max int32
	var count int32
	for _, size := range candles {
		if size == max {
			count++
		} else if size > max {
			max = size
			count = 1
		}
	}
	return count
}
