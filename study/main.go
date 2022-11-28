package main

import (
	"fmt"
	"go-study/study/bsort"
	"sort"
)

func main() {
	// 冒泡排序
	fmt.Println("冒泡排序")
	slice := []int{2, 1, 10, 8, 30}
	bsort.BulleSort(slice)
	fmt.Println(slice)
	sortSlice := sort.IntSlice([]int{2, 1, 10, 8, 30})
	bsort.BubbleSortUsingSortPackage(sortSlice)
	fmt.Println(sortSlice)

	// 插入排序
	fmt.Println("插入排序")
	slice = []int{2, 1, 10, 8, 30}
	bsort.InsertSort(slice)
	fmt.Println(slice)
	sortSlice = sort.IntSlice([]int{2, 1, 10, 8, 30})
	bsort.InsertSortUsingSortPackage(sortSlice)
	fmt.Println(sortSlice)

	// 简单选择排序
	fmt.Println("简单选择排序")
	slice = []int{2, 1, 10, 8, 30}
	bsort.SelectSort(slice)
	fmt.Println(slice)
	sortSlice = sort.IntSlice([]int{2, 1, 10, 8, 30})
	bsort.SelectSortUsingSortPackage(sortSlice)
	fmt.Println(sortSlice)

	// 快速排序
	fmt.Println("快速排序")
	slice = []int{2, 1, 10, 8, 30}
	bsort.QuickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)

	// 快速排序
	fmt.Println("归并排序")
	slice = []int{2, 1, 10, 8, 30}
	slice = bsort.MergeSort(slice)
	fmt.Println(slice)
}
