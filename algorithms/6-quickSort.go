package algo

import "fmt"

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	index := low - 1

	for i := low; i < high; i++ {
		if arr[i] < pivot {
			index++
			temp := arr[i]
			arr[i] = arr[index]
			arr[index] = temp
		}
	}
	index++
	temp := arr[index]
	arr[index] = arr[high]
	arr[high] = temp

	return index
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		index := partition(arr, low, high)
		quickSort(arr, low, index-1)
		quickSort(arr, index+1, high)
	}
}

func ExecuteQuickSort() {
	arr := [6]int{8, 7, 0, 2, 13, 78}
	quickSort(arr[:], 0, len(arr)-1)
	fmt.Println(arr)
}
