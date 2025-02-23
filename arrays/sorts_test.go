package arrays_test

import (
	"reflect"
	"testing"

	"github.com/BambooRaptor/dsa-training/arrays"
)

var tests = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name:     "Empty slice",
		input:    []int{},
		expected: []int{},
	},
	{
		name:     "Single element",
		input:    []int{5},
		expected: []int{5},
	},
	{
		name:     "Already sorted",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Reverse sorted",
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Unsorted with duplicates",
		input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
		expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
	},
	{
		name:     "Negative numbers",
		input:    []int{-5, -1, -3, 0, 2},
		expected: []int{-5, -3, -1, 0, 2},
	},
	{
		name:     "Mixed positive and negative",
		input:    []int{3, -2, 0, -5, 4},
		expected: []int{-5, -2, 0, 3, 4},
	},
}

type Sorter[T any] func(*[]T, arrays.ShouldLRSwap[T])

func testSorts(name string, sorter Sorter[int]) func(*testing.T) {
	return func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				arr := make([]int, len(tt.input))
				copy(arr, tt.input)

				sorter(&arr, arrays.SortIntAsc)

				if !reflect.DeepEqual(arr, tt.expected) {
					t.Errorf("\n%v(%v) => %v\n[Expected] %v\n", name, tt.input, arr, tt.expected)
				}
			})
		}
	}

}

func TestSorts(t *testing.T) {
	t.Run("SortBubble", testSorts("Bubble", arrays.SortBubble))
	t.Run("SortInsertion", testSorts("Insertion", arrays.SortInsertion))
}
