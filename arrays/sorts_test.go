package arrays_test

import (
	"reflect"
	"testing"

	"github.com/BambooRaptor/dsa-training/arrays"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			arrays.SortBubble(&arr, arrays.SortIntAsc)

			if !reflect.DeepEqual(arr, tt.expected) {
				t.Errorf("\nSortBubble(%v) => %v\n[Expected] %v\n", tt.input, arr, tt.expected)
			}
		})
	}
}
