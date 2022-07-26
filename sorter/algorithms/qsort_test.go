package algorithms

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testcases := [][]int{
		{7, 3, 1, 2, 3, 11, 9, 23, 5, 4, 6},
		{1, 2, 3, 4, 5},
		{3, 1},
		{5},
		{},
	}
	excepts := [][]int{
		{1, 2, 3, 3, 4, 5, 6, 7, 9, 11, 23},
		{1, 2, 3, 4, 5},
		{1, 3},
		{5},
		{},
	}

	for i, testcase := range testcases {
		arr := make([]int, len(testcase))
		copy(arr, testcase)
		QuickSort(arr)
		if !reflect.DeepEqual(arr, excepts[i]) {
			t.Errorf("TestCase %v didn't pass. Except %v. Output %v.", testcase, excepts[i], arr)
		}
	}
}
