package main

type SortStrategy interface {
	Sort([]int) []int
}
