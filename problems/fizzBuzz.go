package problems

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Given an integer n, return a string array answer (1-indexed) where:

answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.

Example 1:

Input: n = 3
Output: ["1","2","Fizz"]
Example 2:

Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]
Example 3:

Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]

*/

func FizzBuzz(n int) []string {
	if n < 0 {
		return []string{}
	}
	strArr := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		var sb strings.Builder
		if i%3 == 0 {
			sb.WriteString("Fizz")
		}
		if i%5 == 0 {
			sb.WriteString("Buzz")
		}
		if sb.Len() == 0 {
			sb.WriteString(strconv.Itoa(i))
		}
		strArr = append(strArr, sb.String())
	}
	return strArr
}

func FizzBuzzClassic(n int) []string {
	strArr := []string{}
	for i := 0 + 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			strArr = append(strArr, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			strArr = append(strArr, "Fizz")
			continue
		}
		if i%5 == 0 {
			strArr = append(strArr, "Buzz")
			continue
		}
		strNum := strconv.Itoa(i)
		strArr = append(strArr, strNum)
	}
	fmt.Println("My string array", strArr)
	return strArr
}
