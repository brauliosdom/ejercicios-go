package main

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var count int32
	for i := 0; i < len(ar); i++ {
		count += ar[i]
	}

	return count
}
