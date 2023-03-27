package main

import (
	"strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	n := len(s)
	hour, _ := strconv.Atoi(s[0:2])
	if s[n-2:n] == "AM" {
		if hour == 12 {
			return "00" + s[2:n-2]
		}
		return s[:n-2]
	}
	hour += 12
	if hour == 24 {
		return "12" + s[2:n-2]
	}
	return strconv.Itoa(hour) + s[2:n-2]
}

// testing
func main() {
	println(timeConversion("12:45:54PM"))
}
