package main

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	res := []int32{0, 0}
	for i := 0; i < len(arr); i++ {
		res[0] += arr[i][i]
		res[1] += arr[len(arr)-1-i][i]
	}
	aux := res[0] - res[1]
	if aux < 0 {
		return -1 * aux
	}
	return aux
}

// testing
func main() {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	println(diagonalDifference(arr))
}
