package algo

import "fmt"

func BubbleSortWithPointer(arr []int) {

	distanceToTravel := len(arr)

	pointer1 := 0
	pointer2 := 1

	for distanceToTravel >= 2 {
		if arr[pointer2] < arr[pointer1] {
			// we swap the places.
			temp := arr[pointer1]
			arr[pointer1] = arr[pointer2]
			arr[pointer2] = temp
		}
		pointer1 += 1
		pointer2 += 1

		if pointer2 == distanceToTravel {
			// once the pointer reach the end. Reduce the distance to travel.
			distanceToTravel -= 1
			// then, reset the loop.
			pointer1 = 0
			pointer2 = 1
		}
	}
}

func bubbleSortWithOnlyLoops(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func ExecuteBubbleSort() {
	array := [7]int{9, 3, 7, 4, 69, 420, 42}
	sortedArray := []int{3, 4, 7, 9, 42, 69, 420}

	// bubbleSortWithPointer(array[:])
	bubbleSortWithOnlyLoops(array[:])
	if array == [7]int(sortedArray) {
		fmt.Println(true)
	}
}
