package main

import "fmt"

//This strategy pattern is used to sort arrays

func main() {
	arr := []int{7, 6, 5, 4, 3, 2, 1}

	bubble_sort := BubbleSortStrategy{}
	quick_sort := QuickSortStrategy{}

	context := NewSortingContext(bubble_sort)
	sortedArr := context.Sort(arr)
	fmt.Println("Sorted by Bubble Sort: ", sortedArr)

	context.SetStrategy(quick_sort)
	sortedArr = context.Sort(arr)
	fmt.Println("Sorted by Quick Sort: ", sortedArr)
}
