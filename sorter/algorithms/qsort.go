package algorithms

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivot, i, j := arr[0], 0, len(arr)-1
	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[0], arr[j] = arr[j], arr[0]
	QuickSort(arr[:j])
	QuickSort(arr[j+1:])
}
