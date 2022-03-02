package sort

func partition(arr []int, low, high int) int {
	index := low - 1
	pivotEle := arr[high]
	for i := low; i < high; i++ {
		if arr[i] <= pivotEle {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]
		}
	}

	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

func QuickSortRange(arr []int, low, high int) {
	if len(arr) <= 1 {
		return
	}

	if low < high {
		pivot := partition(arr, low, high)
		QuickSortRange(arr, low, pivot-1)
		QuickSortRange(arr, pivot+1, high)
	}
}

func QuickSort(arr []int) []int {
	QuickSortRange(arr, 0, len(arr)-1)
	return arr
}
