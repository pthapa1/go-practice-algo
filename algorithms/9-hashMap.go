package algo

import "strconv"

/*
Hash table is basically an array where we use hash function to generate hash index,
and we use the generated hash index to insert and search data.

Since we  are getting the item with index
- searching, inserting and deleting items are all constant time.
- Worst case scenario is O(n) for all of them
- Listing all the items is O(n)

***
components
key: Can be anything, an integer or string. We put this key in the hash function to generate hash index
Hash Function: Function that takes in key and generates the hash inxex.
Hash Table: Hash table is a data structure that maps the key

***
issue with hash table
The problem is our hash function could generate the same hash index for two different keys.
When that happens, we need to find ways to store the data without replacing what already exists in the table

👆: This is called collision. To solve it, we need to handle collision.
- Open Addressing: We could either place the item in the cell next to the item with the same index
  - Searching: Go to the original place, if it's not there then move forward (liner)

  - As the array gets larger, we get farther away from the original hash index, and it becomes a liner array

- **Separate Chaining**: We could place the item one after the other in the same index but with linked list attached to it.

  - Instead of putting the item in the array, each index points to the head of a linked list

    0  1  2  3  4  5  6  7  <- index
  |  |  |  |  |  |  |  |  | <- array (associative) each slot is bucket that points to head of the linked list
	 ⬇  ⬇  ⬇  ⬇  ⬇  ⬇  ⬇  ⬇   <- each bucker points to linked list (bucker nodes)
	___
	___   <- bucket nodes

***
Rules in separate chaining:
- We store any key in the bucket, linked list
- All keys have to be unique. That means, we need to check whether that key already exists on the bucket

***
Insert:
- Generate hash code using hash function
- Insert the item in the provided index by hash function

***
Search:
- Generate hash code for the hash function again,
- Go to the given index to search for the item
*/
// --- Separate chaining ---

// Takes either string or an int, and an array size, an int.
// And returns a unique (index) for the key
func Hash[T int | string](key T, arraySize int) int {
	sum := 0
	var convertedValue string
	if v, ok := any(key).(int); ok {
		convertedValue = strconv.Itoa(v)
	} else if v, ok := any(key).(string); ok {
		convertedValue = v
	}
	for _, val := range convertedValue {
		sum += int(val)
	}
	return sum % arraySize
}

type HashTable[T int | string] struct {
	array []*HashNode[T]
}

type HashNode[T int | string] struct {
	Data T
	Next *HashNode[T]
}

// Initialize & Returns an array of arraySize, where each item is of type HashNode
func Init[T int | string](arraySize int) *HashTable[T] {
	result := &HashTable[T]{
		array: make([]*HashNode[T], arraySize),
	}
	for i := range result.array {
		result.array[i] = &HashNode[T]{}
	}
	return result
}

/*
- Need an item to insert (key/ key-value)
- Use hash function to generate the hash key.
- Use the hash key to insert the item in the hash table.
  - First, you need to check if the item exists in the hashmap
  - if the item exists, we update it at that location. If it does not exist,
    - Save the old head
    - Add the data in the head.
    - Point the head to the new data
    - Point the new head to the old head
** This assumes that we don't allow duplicate keys
*/
// Insert item to hashmap
func (h *HashTable[T]) HashMapInsert(key T) {
	idx := Hash[T](key, len(h.array))
	/* make a copy of current node we are trying to access and check if it's nil.
	if the copy's value is nil, that means the place we are trying to put the items in is also nil */
	currentNode := h.array[idx]
	if currentNode == nil {
		h.array[idx] = &HashNode[T]{Data: key, Next: nil}
	} else {
		// when we end up at the end of the list, we need to make sure we have saved the last item
		// so we can insert the item at the end
		var lastNode *HashNode[T]
		// if it is not nil, (or node exists) check if the item already exists
		for currentNode != nil {
			if currentNode.Data == key {
				/* In this case, we don't have a map with key value pair, so we can simply return but
				keeping the currentNode.Data anyway*/
				currentNode.Data = key
				return
			}
			lastNode = currentNode
			currentNode = currentNode.Next
		}
		/* We traversed the entire list but the key does not exist. So, let's add this unique item at the end.
		We can also add this item at the head. In that case we need to save the initial head instead
		*/
		lastNode.Next = &HashNode[T]{Data: key, Next: nil}
	}
}

/*
- Need an item to Search (key)
- Generate Hash idx with hash function
- Search the entire list
- if found, return true, in our case we don't have value to return
- if not, return false
*/

// search if key exists in a hashmap
func (h *HashTable[T]) HashMapSearch(key T) bool {
	idx := Hash[T](key, len(h.array))
	currentNode := h.array[idx]
	for currentNode != nil {
		if currentNode.Data == key {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

// delete node from hashmap given a key

func (h *HashTable[T]) HashMapDelete(key T) bool {
	idx := Hash[T](key, len(h.array))
	currentNode := h.array[idx]
	var lastNode *HashNode[T] = nil

	// Check if the head node itself holds the key to be deleted
	if currentNode != nil && currentNode.Data == key {
		// Move the head to the next node
		h.array[idx] = currentNode.Next
		return true
	}

	// Search for the key to be deleted, keep track of the previous node
	// as we need to change 'next' of the previous node
	for currentNode != nil {
		if currentNode.Data == key {
			// Unlink the node from linked list
			if lastNode != nil {
				lastNode.Next = currentNode.Next
			}
			return true
		}
		lastNode = currentNode
		currentNode = currentNode.Next
	}

	// If the key was not present in linked list
	return false
}

/*
// we can also create a distinct linked list at each array's space.
// in that case, we need to make the following changes
type HashLinkedList[T int | string] struct {
	Head   *HashNode[T]
	Length int
}

type HashTable[T int | string] struct {
    array []*HashLinkedList[T] // Use HashLinkedList instead of HashNode
}

// Revised Init function to use HashLinkedList
func Init[T int | string](arraySize int) *HashTable[T] {
    result := &HashTable[T]{
        array: make([]*HashLinkedList[T], arraySize),
    }
    // Initialize each entry with an empty HashLinkedList
    for i := range result.array {
        result.array[i] = &HashLinkedList[T]{Head: nil, Length: 0}
    }
    return result
}

// Modified HashMapInsert to use HashLinkedList
func (h *HashTable[T]) HashMapInsert(key T) {
    idx := Hash[T](key, len(h.array))
    // Access the HashLinkedList at the index
    list := h.array[idx]
    if list.Head == nil {
        list.Head = &HashNode[T]{Data: key, Next: nil}
        list.Length = 1
    } else {
        // Insert at the head for simplicity
        newNode := &HashNode[T]{Data: key, Next: list.Head}
        list.Head = newNode
        list.Length++
    }
}
*/
