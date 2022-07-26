package algorithms

func merge(arr []int, mid int) {
	part1 := make([]int, mid)
	copy(part1, arr)
	part2 := arr[mid:]
	i, i1, i2 := 0, 0, 0

	for i1 < len(part1) && i2 < len(part2) {
		if part1[i1] <= part2[i2] {
			arr[i] = part1[i1]
			i, i1 = i+1, i1+1
		} else {
			arr[i] = part2[i2]
			i, i2 = i+1, i2+1
		}
	}
	for i1 < len(part1) {
		arr[i] = part1[i1]
		i, i1 = i+1, i1+1
	}
}

func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	mid := len(arr) / 2
	MergeSort(arr[:mid])
	MergeSort(arr[mid:])
	merge(arr, mid)
}
