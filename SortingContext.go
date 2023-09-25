package main

type SortingContext struct {
	strategy SortStrategy
}

func NewSortingContext(strategy SortStrategy) *SortingContext {
	return &SortingContext{strategy: strategy}
}

func (sortingcontext *SortingContext) SetStrategy(strategy SortStrategy) {
	sortingcontext.strategy = strategy
}

func (sortingcontext *SortingContext) Sort(arr []int) []int {
	return sortingcontext.strategy.Sort(arr)
}
