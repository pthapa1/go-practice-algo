package main

import "fmt"

func mapHasItem(obj map[string]bool, item string) bool {
	return obj[item]
}

func main() {
	input := make(map[string]bool)
	input["1"] = true
	result := mapHasItem(input, "2")
	fmt.Println(result)
	i, v := input["1"]
	fmt.Println("This is true, first value: ", i)

	fmt.Println("This is false, second value: ", v)
}
