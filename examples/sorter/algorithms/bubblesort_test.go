package algorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testcases := [][]int{
		{7, 3, 1, 2, 3, 5, 4, 6},
		{1, 2, 3, 4, 5},
		{3, 1},
		{5},
		{},
	}
	excepts := [][]int{
		{1, 2, 3, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 5},
		{1, 3},
		{5},
		{},
	}

	for i, testcase := range testcases {
		arr := make([]int, len(testcase))
		copy(arr, testcase)
		BubbleSort(arr)
		if !reflect.DeepEqual(arr, excepts[i]) {
			t.Errorf("TestCase %v didn't pass. Except %v. Output %v.", testcase, excepts[i], arr)
		}
	}
}
