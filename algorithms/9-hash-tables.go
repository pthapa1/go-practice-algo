package algo

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
// hash function: takes key, of string, and an array size, an int. And returns a unique (index) for the key
func Hash(key string, arraySize int) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize
}

type HashTable[T int | string] struct {
	array []*HashNode[T] // array of size ArraySize and type bucket
}

type HashNode[T int | string] struct {
	Data T
	Next *Node[T]
}

type HashLinkedList[T int | string] struct {
	Head   *HashNode[T]
	Length int
}

func Init[T int | string](arraySize int) *HashTable[T] {
	result := &HashTable[T]{
		array: make([]*HashNode[T], arraySize),
	}
	for i := range result.array {
		result.array[i] = &HashNode[T]{}
	}
	return result
}

//
// // bucket structure
// type bucket[T int | string] struct {
// 	head *HashLinkedListNode[T]
// }
//
// // bucketNode is a linked list, so we need to define the individual node structure as well
// type bucketNode struct {
// 	next *bucketNode
// 	key  string
// }

// Initialize a bucket: In each slot of the HashTable array add a bucket structure
// func Init() *HashTable {
// 	result := &HashTable{}
// 	for i := range result.array {
// 		result.array[i] = &bucket{}
// 	}
// 	return result
// }

// insert
// search
// delete

// insert
// search
// delete

// hash function

// init
