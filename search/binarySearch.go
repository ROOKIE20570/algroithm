package main

import "fmt"

func main() {
	arr := []int{2,5,7,7,9,10}
	res := findLast(arr,7)
	fmt.Println(res)
}

func search(arr []int, val int) int {
	length := len(arr)
	low := 0
	high := length - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1

}

func findLast(arr []int, val int) int {
	length := len(arr)
	low := 0
	high := length - 1
	
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == val {
			if mid == 0||  arr[mid + 1] != val{
				return mid
			}else{
				low = mid + 1
			}
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

