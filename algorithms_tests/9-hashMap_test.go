package algo

import (
	"testing"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func TestHashMap(t *testing.T) {
	testIntCases := []struct {
		name      string
		key       []int
		arraySize int
	}{
		{
			name:      "Int",
			key:       []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9},
			arraySize: 10,
		},
		{
			name:      "Int2",
			key:       []int{22, 33, 44, 55, 66, 77, 88, 99, 100},
			arraySize: 4,
		},
	}
	testStringCases := []struct {
		name      string
		key       []string
		arraySize int
	}{
		{
			name:      "String",
			arraySize: 4,
			key: []string{
				"one",
				"two",
				"three",
				"four",
				"five",
				"six",
				"seven",
				"eight",
				"nine",
			},
		},
		{
			name:      "String2",
			arraySize: 9,
			key: []string{
				"berry",
				"bary",
				"berries",
				"bury",
				"buried",
			},
		},
	}

	for _, caseInt := range testIntCases {
		t.Run(caseInt.name, func(t *testing.T) {
			hashTable := algo.Init[int](caseInt.arraySize)
			for _, key := range caseInt.key {
				hashTable.HashMapInsert(key)
			}
			for _, key := range caseInt.key {
				found := hashTable.HashMapSearch(key)
				if !found {
					t.Errorf("key %v not found", key)
				}
			}
		})
	}

	for _, caseStr := range testStringCases {
		t.Run(caseStr.name, func(t *testing.T) {
			hashTable := algo.Init[string](caseStr.arraySize)
			for _, key := range caseStr.key {
				hashTable.HashMapInsert(key)
			}
			for _, key := range caseStr.key {
				found := hashTable.HashMapSearch(key)
				if !found {
					t.Errorf("key %v not found", key)
				}
			}
		})
	}
}
