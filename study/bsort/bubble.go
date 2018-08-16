package bsort

import "sort"

/**
* 使用golang实现常见排序算法，以及使用sort包来实现这些算法：统一都是升序排序
*/

// 冒泡排序
func BulleSort(items []int) {
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items) - i - 1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}

// 冒泡排序 使用sort包
func BubbleSortUsingSortPackage(data sort.Interface) {
	len := data.Len()
	for i := 0; i < len - 1; i++ {
		for j := 0; j < len - 1 - i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
			}
		}
	}
}

// 插入排序
func InsertSort(items []int) {
	for i := 1; i < len(items); i++ {
		for j := i; j > 0; j-- {
			if items[j-1] <= items[j] {
				// 已经是有序，不需要再继续比较
				break
			}

			items[j-1], items[j] = items[j], items[j-1]
		}
	}
}

// 插入排序 使用sort包
func InsertSortUsingSortPackage(data sort.Interface) {
	r := data.Len() - 1
	for i := 1; i <= r; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// 简单选择排序
func SelectSort(items []int) {
	for i := 0; i < len(items) - 1; i++ {
		var minIdx = i
		for j := i+1; j < len(items); j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

// 简单选择排序 使用sort包
func SelectSortUsingSortPackage(data sort.Interface) {
	r := data.Len() - 1
	for i := 0; i < r; i++ {
		min := i
		for j := i + 1; j <= r; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}
