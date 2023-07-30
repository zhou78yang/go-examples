package algorithms

import "testing"

func TestCheckOrder(t *testing.T) {
	testcases := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 4, 3, 5},
		{},
		{4},
		{2, 1},
	}
	expects := []bool{true, false, true, true, false}

	for i, testcase := range testcases {
		result := CheckOrder(testcase)
		if result != expects[i] {
			t.Errorf("TestCase %v didn't pass. Expect %v . Output %v .", testcase, expects[i], result)
		}
	}
}
