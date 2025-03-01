package method

// sol1. iterative (bottom-up)
func MergeSortIterative(arr []int) {
	for size := 1; size <= len(arr); size *= 2 {
		for left := 0; left < len(arr)-size; left += 2 * size {
			mid := left + size - 1
			right := min(left+2*size-1, len(arr)-1)
			merge2(arr, left, mid, right)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge2(arr []int, left, mid, right int) {
	temp := make([]int, len(arr))
	i, j, k := left, mid+1, left

	// 오름차순 정렬
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}

	// 길이가 2^n 이 아닌 배열 대응
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	// 원래 배열에 복사
	for i := left; i <= right; i++ {
		arr[i] = temp[i]
	}
}
