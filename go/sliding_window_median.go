package main

import (
	"fmt"
	"sort"
)

// question: https://leetcode.com/problems/sliding-window-median/
func medianSlidingWindow(nums []int, k int) []float64 {
	res := make([]float64, 0)

	window := make([]int, k)
	copy(window, nums[:k])
	sort.Ints(window)
	res = append(res, median(window))
	for i := 0; i < len(nums)-k; i++ {
		window = removeValue(window, nums[i])
		add := nums[k+i]
		index := 0
		for index < len(window) && window[index] < add {
			index++
		}
		window = insertAt(window, index, add)
		res = append(res, median(window))
	}

	return res
}

func median(nums []int) float64 {
	if len(nums)%2 == 0 {
		return float64((nums[len(nums)/2] + nums[len(nums)/2-1])) / 2.0
	}

	return float64(nums[len(nums)/2])
}

// 0 <= index <= len(a)
func insertAt(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func removeValue(slice []int, s int) []int {
	index := 0
	for i, val := range slice {
		if val == s {
			index = i
		}
	}

	if index == len(slice)-1 {
		return slice[:index]
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(medianSlidingWindow([]int{1, 4, 2, 3}, 4))
	// input := []int{-1, 1, 3}
	// fmt.Println(removeValue(input, 1))
}
