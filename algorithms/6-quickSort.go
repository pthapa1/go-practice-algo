package algo

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

func ExecuteQuickSort(arr []int) {
	quickSort(arr[:], 0, len(arr)-1)
}
