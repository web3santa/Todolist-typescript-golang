package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, 2, 1, 5, 3}
	fmt.Println(secondLargest(arr)) // Output: 4
}

func secondLargest(nums []int) int {
	// 배열을 정렬하여 두 번째로 큰 값을 찾음
	sort.Ints(nums)
	fmt.Println(nums)
	n := len(nums)
	return nums[n-2]
}
