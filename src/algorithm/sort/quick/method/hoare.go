package method

func QuickSortHoare(arr []int, low, high int) {
	if low < high {
		pivot := partitionHoare(arr, low, high)
		QuickSortHoare(arr, low, pivot)
		QuickSortHoare(arr, pivot+1, high)
	}
}

func partitionHoare(arr []int, low, high int) int {
	pivot := arr[low]
	left, right := low, high

	for {
		for arr[left] < pivot {
			left++
		}

		for arr[right] > pivot {
			right--
		}

		if left >= right {
			return right
		}

		arr[left], arr[right] = arr[right], arr[left]
	}
}
