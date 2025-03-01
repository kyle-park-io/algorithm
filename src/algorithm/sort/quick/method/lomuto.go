package method

func QuickSortLomuto(arr []int, low, high int) {
	if low < high {
		pivot := partitionLomuto(arr, low, high)
		QuickSortLomuto(arr, low, pivot-1)
		QuickSortLomuto(arr, pivot+1, high)
	}
}

func partitionLomuto(arr []int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}
