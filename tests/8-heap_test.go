package algo

import (
	"reflect"
	"testing"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func TestHeapInsert(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput []int
		head           int
	}{
		{[]int{10, 20, 30, 40, 16, 27}, []int{40, 30, 27, 10, 16, 20}, 40},
		{[]int{1, 5, 7, 8, 9, 14, 16, 30, 34, 50}, []int{50, 34, 14, 16, 30, 5, 9, 1, 8, 7}, 50},
	}
	result := &algo.MaxHeap{}
	for _, tc := range testCases {
		for _, v := range tc.input {
			result.Insert(v)
		}
		if reflect.DeepEqual(tc.expectedOutput, result) {
			t.Errorf(
				"Error on Inserting element on a heap. Expected: %v \n Got: %v",
				tc.expectedOutput,
				result,
			)
		}
		if result.Slice[0] != tc.head {
			t.Errorf(
				"Maxheap's head should be the largest in the slice. \n Expected %v \n Got %v",
				tc.head,
				result.Slice[0],
			)
		}
	}
}

func TestHeapExtraction(t *testing.T) {
	testCases := []struct {
		input            []int
		expectedHeap     []int
		expectedMax      int
		expectedAfterMax []int // expected heap structure after extracting the max
	}{
		{
			input:            []int{10, 20, 30, 40, 16, 27},
			expectedHeap:     []int{40, 30, 27, 10, 16, 20},
			expectedMax:      40,
			expectedAfterMax: []int{30, 20, 27, 10, 16},
		},
		{
			input:            []int{1, 5, 7, 8, 9, 14, 16, 30, 34, 50},
			expectedHeap:     []int{50, 34, 14, 16, 30, 5, 9, 1, 8, 7},
			expectedMax:      50,
			expectedAfterMax: []int{34, 30, 14, 16, 7, 5, 9, 1, 8},
		},
	}

	for _, tc := range testCases {
		// Build the heap
		heap := &algo.MaxHeap{}
		for _, v := range tc.input {
			heap.Insert(v)
		}

		// Test the heap structure
		if !reflect.DeepEqual(heap.Slice, tc.expectedHeap) {
			t.Errorf("Heap structure incorrect. Expected: %v, got: %v", tc.expectedHeap, heap.Slice)
		}

		// Test extracting the maximum element
		extractedMax := heap.ExtractLargestVal()
		if extractedMax != tc.expectedMax {
			t.Errorf(
				"Extracted max element incorrect. Expected: %d, got: %d",
				tc.expectedMax,
				extractedMax,
			)
		}

		// Test the heap structure after extraction
		if !reflect.DeepEqual(heap.Slice, tc.expectedAfterMax) {
			t.Errorf(
				"Heap structure after extraction incorrect. Expected: %v, got: %v",
				tc.expectedAfterMax,
				heap.Slice,
			)
		}
	}
}
