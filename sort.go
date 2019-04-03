package sort

func bubblesort(a []int) []int {

	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
	}
	return a
}

func quicksort(a []int) []int {

	l, r := 0, len(a)-1

	if l < r {
		for i := range a {
			if a[i] < a[r] {
				a[l], a[i] = a[i], a[l]
				l++
			}
		}
		a[l], a[r] = a[r], a[l]
		quicksort(a[:l])   // sort leftmost values
		quicksort(a[l+1:]) // soft rightmost value
	}
	return a
}

func mergesort(a []int) []int {

	if len(a) < 2 {
		return a
	}

	m := len(a) / 2
	return mergeslice(mergesort(a[:m]), mergesort(a[m:]))
}

// merges two int-slices, returning a sorted slice
func mergeslice(a, b []int) []int {

	size, i, j := len(a)+len(b), 0, 0
	res := make([]int, size)

	for k := 0; k < size; k++ {

		if i > len(a)-1 && j <= len(b)-1 {
			// if end of "a" has been reached, use b
			res[k] = b[j]
			j++
		} else if i <= len(a)-1 && j > len(b)-1 {
			// if end of "b" has been reached, use a
			res[k] = a[i]
			i++
		} else if a[i] < b[j] {
			// if no array limits have been met, simply use the greatest value
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}

	}

	return res
}
