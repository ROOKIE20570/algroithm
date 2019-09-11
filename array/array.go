package main

import "fmt"

type myArray []int

func (arr myArray) unique() []int {
	tmpMap := make(map[int]bool)
	var newArr myArray
	for i := 0; i < len(arr); i++ {
		if _,ok := tmpMap[arr[i]];!ok {
			newArr = append(newArr,arr[i])
			tmpMap[arr[i]] = true
		}
	}

	return newArr
}

func main() {
	a := myArray{1, 11, 5, 6, 6, 8, 11}

	fmt.Println(a.unique())
}

func mergeSortedArr(arr1, arr2 myArray) myArray {
	if len(arr1) == 0 {
		return arr2
	}

	if len(arr2) == 0 {
		return arr1
	}

	var tmpArr myArray

	arr1Cur := 0
	arr2Cur := 0
	for i := 0; i < len(arr1)+len(arr2); i++ {
		if arr1[arr1Cur] < arr2[arr2Cur] {
			tmpArr = append(tmpArr, arr1[arr1Cur])
			if arr1Cur+1 == len(arr1) {
				break
			}
			arr1Cur++
		} else {
			tmpArr = append(tmpArr, arr2[arr2Cur])
			if arr2Cur+1 == len(arr2) {
				break
			}
			arr2Cur++
		}

	}

	return tmpArr
}
