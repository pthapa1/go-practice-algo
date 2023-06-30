package algo

import "fmt"

func basicExample(number int) {
	if number <= 0 {
		fmt.Println("Done!")
		return
	}
	fmt.Println(number)
	basicExample(number - 1)
}

func basicExample1(number int) {

	newNumber := number - 1

	if number >= 1 {
		fmt.Println(number)
		basicExample1(newNumber)
	}
}

func basicExample2(number int) int {
	if number <= 1 {
		fmt.Println(1)
		return 1
	}
	result := basicExample2(number-1) + number
	fmt.Println(result)
	return result
}

func ExecuteRecursion() {
	basicExample1(5)
}
