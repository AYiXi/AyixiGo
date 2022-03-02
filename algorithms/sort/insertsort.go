package sort

func InsertSort(arr []int) []int {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		temp := arr[currentIndex]
		iterator := currentIndex
		for ; iterator > 0 && arr[iterator-1] >= temp; iterator-- {
			arr[iterator] = arr[iterator-1]
		}
		arr[iterator] = temp
	}

	return arr
}
