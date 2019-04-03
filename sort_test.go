package sort

import (
	"testing"
)

func Test_allSorts(t *testing.T) {

	type sortFunc func([]int) []int
	var (
		in   = []int{2, 12, 1, 9, 5}
		want = []int{1, 2, 5, 9, 12}
	)
	testCases := []struct {
		name string
		fn   sortFunc
	}{
		{
			name: "quicksort",
			fn:   quicksort,
		},
		{
			name: "mergesort",
			fn:   mergesort,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.fn(in); !sliceEqual(got, want) {
				t.Errorf("wanted %v, got %v", want, got)
			}
		})
	}
}

// func Test_quicksort(t *testing.T) {
// 	unsorted := []int{2, 12, 1, 9, 5}
// 	got := quicksort(unsorted)
// 	want := []int{1, 2, 5, 9, 12}
// 	if !sliceEqual(got, want) {
// 		t.Errorf("wanted %v, got %v", want, got)
// 	}
// }

// func Test_mergesort(t *testing.T) {
// 	unsorted := []int{2, 12, 1, 9, 5}
// 	got := mergesort(unsorted)
// 	want := []int{1, 2, 5, 9, 12}
// 	if !sliceEqual(got, want) {
// 		t.Errorf("wanted %v, got %v", want, got)
// 	}
// }

// bubblesort
// heapsort
// insertionsort

func sliceEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
