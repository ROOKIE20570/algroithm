package main

import (
	"sync"
)

var mu sync.Mutex
func main() {


}

//查找数组中重复的数字
//在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。请找出数组中任意一个重复的数字。
func findUnique(arr []int) bool {
	if arr == nil {
		return false
	}

	for i := 0; i < len(arr); i++ {
		for i != arr[i] {
			if arr[i] != arr[arr[i]] {
				return true
			}

			swap(arr, i, arr[i])
		}
	}
	return false
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
