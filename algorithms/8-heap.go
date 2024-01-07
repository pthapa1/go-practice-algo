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

Importnat formula ‼️
indexOfLeftChild = parentIndex * 2 + 1
indexofRightChild = parentIndex * 2 + 2 (or add 1 to left child)

--- Insert
To insert a node to the tree,
* Put the item at the end of the array. (Viaually: end open position on the tree above)
* Then, we compare the inserted number with it's parent, with formulas above.
  * if the parent is smaller, we swap the position with it's parent.
  * Use a while loop until the parent is the largest.

--- Extraction: highest key
To extract the highest key
* extract the first item in the array. (Viaually: the root of the tree)
* Copy the value of the last item in the array to the head (can't leave a tree empty)
* Remove the last item from the tree, since we just copied it's value to the root
* finally, heapify down, compare it with it's largest children, if any child is larger than the val, swap
>  Time Complexity O(log n) -- n is the number of items in the array
>  Other time complexity of the heap depends on the worst part about this data structure.

In heap, the worst part about the operation is calculating
the place for parent/child (traversing the tree up and down) when an item is
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

	// Continue heapifying down as long as the right child of the current parent
	// is within the bounds of the heap.
	for right(index) <= lastIndex {
		l, r := left(index), right(index)
		largest := index

		// Check if left child exists and is greater than the parent
		if l <= lastIndex && h.Slice[l] > h.Slice[largest] {
			largest = l
		}

		// Check if right child exists and is greater than the current largest
		if r <= lastIndex && h.Slice[r] > h.Slice[largest] {
			largest = r
		}

		// If the largest is still the parent, the heap property is satisfied
		if largest == index {
			break
		}

		// Swap the parent with the largest child
		h.swap(index, largest)

		// Move down the tree
		index = largest
	}
	/*
		// or we dan do this too
				 // while at least 1 left child exists that's smaller or equal than the lastIndex
				 for left(index) <= lastIndex {
				 	l, r := left(index), right(index)
				 	largest := l

					if r <= lastIndex && h.Slice[r] > h.Slice[l] {
					largest = r
				 	}

					// we are done if parent is larger than the largest child
					if h.Slice[index] >= h.Slice[largest] {
					break
					}
					h.swap(index, largest)
					// update the index so the while loop can continue
					index = largest

					 }
	*/
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
	// remove the last item. Make slice shorter
	h.Slice = h.Slice[:lastIndex]
	h.maxHeapifyDown(0)
	return extractedValue
}
