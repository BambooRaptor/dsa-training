package arrays

func swap[T any](arr *[]T, i int, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}

func SortIntAsc(a int, b int) bool {
	return a > b
}

// slrs => should left <=> right swap?
func SortBubble[T comparable](arr *[]T, slrs func(T, T) bool) {
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
