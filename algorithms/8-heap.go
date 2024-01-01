/*
* Heap is a data structure mainly used for priority queue.
* In normal queue, we take out the one that came in first,
but in priority queue, we take out what's important first.
* Also, heap can be expressed as a *complete tree*. A cmolplete tree,
  is a tree where all the nodes are full except for the leaves (the lowest level)
  where all the filled nodes are to the left.
           50
         /    \
       30      40
      /  \    /  \
    20   15  35  25
   /  \
  10  5

* Max heap 👆, where each parent key is larger than it's children
* Min heap, where each parent has the smallest of the keys. Just opposite of the max heap
---

* Heap is illustrated and operates as a tree but is acatually an array under the hood.
for the tree above, the array would look like this
heapArray := []int{50, 30, 40, 20, 15, 35, 25, 10, 5}
* To calculate the index of children based on the parent's index,

indexOfLeftChild = parentIndex * 2 + 1
indexofRightChild = parentIndex * 2 + 2 (or add 1 to left child)
👆 this is what we will use to calculate and find the child and parents

--- Insert
To insert a node to the tree,
* first put the node to the bottom right of the tree. The last index in the array
* Then, we compare the inserted number with it's parent.
  * if the parent is smaller, we swap the position with it's parent.
  * recurse the step until the parent is larger then the inserted item.

--- Extraction
To extract the highest key extract the root
* first replace the root with the last item in the array. In our example, we would replace 50 with 5.
* Compare the new root (5) with it's children, if any of the child is larger than the root, swap places.
* recurse until the node's children are both smaller than the parent

--- Time Complexity O(log n) -- n is the number of items in the array
Time complexity of the heap depends on the worst part about this data structure.
In heap, the worst part about the operation is traversing the tree up and down when an item is
* extracted
* inserted
In these two operations, the larger the height, the longer time it takes.
As the number of items increase, the height increases but not in a linear fashion.
It increases with log n
*/

// rememnber: in max heap, parent's value is always greater

package algo

import "fmt"

type MaxHeap struct {
	Slice []int
}

// inserting an element to the heap
func (h *MaxHeap) Insert(key int) {
	// first append to the end
	h.Slice = append(h.Slice, key)
	indexOfKey := len(h.Slice) - 1 // since we appended the key as the last element
	h.maxHeapifyUp(indexOfKey)
}

// returns the parent's index, when you provide any child's index
func parent(i int) int {
	// if it's a parent index,
	if i == 0 {
		return 0
	}
	// right child is always an even number
	if i%2 == 0 {
		return (i - 2) / 2
	}
	// Left child is always an odd number
	return (i - 1) / 2
	// in fact this funciton always rounds down,
	// so we can just return (i - 1 ) / 2
	// but for my sanity, I am keeping it verbose
}

// Given two index, swap the actual values those index represent in the slice
func (h *MaxHeap) swap(i1, i2 int) {
	h.Slice[i1], h.Slice[i2] = h.Slice[i2], h.Slice[i1]
}

// maxHeapifyUp() will heapify from bottom up
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.Slice[parent(index)] < h.Slice[index] {
		h.swap(parent(index), index) // swap first
		// when we swap, we need to find the parent of index again, so that, the while loop can continue
		// think of this like a recursion,
		index = parent(index)
	}
}

// given parent index, find the left index
func left(i int) int {
	return (i * 2) + 1
}

// given parent index, find the right index
func right(i int) int {
	return (i * 2) + 2
}

// maxHeapifyDown will heapify from top down
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.Slice) - 1

	for left(index) <= lastIndex {
		l, r := left(index), right(index)
		indexToCompare := l
		// if right index is within slice bounds(exists)
		// and right value is greater than the left value
		if r <= lastIndex && h.Slice[r] > h.Slice[l] {
			indexToCompare = r
		}
		// we are done if parent is larger than the largest child
		if h.Slice[index] >= h.Slice[indexToCompare] {
			break
		}
		h.swap(index, indexToCompare)
		// update the index so the while loop can continue
		index = indexToCompare

	}
}

// extract the largest value
func (h *MaxHeap) ExtractLargestVal() int {
	extractedValue := h.Slice[0] // first value is the largest value
	lastIndex := len(h.Slice) - 1
	if len(h.Slice) == 0 {
		fmt.Println("Cannot extract from empty array")
		return -1
	}
	// get the last item from the slice and copy it's value in head
	h.Slice[0] = h.Slice[lastIndex]
	// remove the last item. Make slisce shorter
	h.Slice = h.Slice[:lastIndex]
	h.maxHeapifyDown(0)
	return extractedValue
}
