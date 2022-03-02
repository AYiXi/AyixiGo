package main

import (
	"ayixigo/algorithms/sort"
	"fmt"
)

func main() {
	// quicksort
	nums := []int{3, 1, 2, 4, 5, 8, 7, 6}
	var r []int

	r = sort.QuickSort(nums)
	fmt.Println("QuickSort:", r)

	r = sort.BubbleSort(nums)
	fmt.Println("BubbleSort:", r)

	r = sort.InsertSort(nums)
	fmt.Println("InsertSort:", r)
}
