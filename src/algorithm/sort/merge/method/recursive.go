package method

// sol1. recursive (top-down)
func MergeSortRecursive(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2

	// warn: overflow
	// mid := (left + right) / 2

	MergeSortRecursive(arr, left, mid)
	MergeSortRecursive(arr, mid+1, right)
	merge(arr, left, mid, right)
}

// mid 가 부분 길이를 나타내는 꼴
func merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

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
	for i, v := range temp {
		arr[left+i] = v
	}
}
