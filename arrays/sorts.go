package arrays

func swap[T any](arr *[]T, i int, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}

func SortIntAsc(a int, b int) bool {
	return a > b
}

type ShouldLRSwap[T any] func(T, T) bool

// slrs => should left <=> right swap?
func SortBubble[T any](arr *[]T, slrs ShouldLRSwap[T]) {
	n := len(*arr)
	for i := range n - 1 {
		swapped := false

		for j := range n - 1 - i {
			shouldSwap := slrs((*arr)[j], (*arr)[j+1])
			if shouldSwap {
				swap(arr, j, j+1)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func SortInsertion[T any](arr *[]T, slrs ShouldLRSwap[T]) {
	n := len(*arr)
	for i := 1; i < n; i++ {
		for j := range i {
			shouldSwap := slrs((*arr)[i-j-1], (*arr)[i-j])
			if !shouldSwap {
				break
			}
			swap(arr, i-j, i-j-1)
		}
	}
}
