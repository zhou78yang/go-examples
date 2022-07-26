package main

import (
	"flag"
	"fmt"

	"sorter/algorithms"
)

var algorithm string

func Input() []int {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	return arr
}

func Output(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func main() {
	flag.StringVar(&algorithm, "a", "qsort", "排序算法")
	flag.Parse()

	arr := Input()

	switch algorithm {
	case "qsort":
		algorithms.QuickSort(arr)
	case "mergesort":
		algorithms.MergeSort(arr)
	case "bubblesort":
		algorithms.BubbleSort(arr)
	default:
		fmt.Println("Unsupport algorithm", algorithm)
	}

	Output(arr)
}
