package main

import "fmt"

type myArray []int

func (arr myArray) unique() []int {
	var a int = 0
	var newArr myArray
	if arr == nil {
		return nil
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr[a] {
			newArr = append(newArr, arr[i])
			a++
		}
	}

	return newArr
}

func main() {
	a := myArray{1, 5, 6, 8, 11}
	b := myArray{1, 6, 7, 7, 9}

	fmt.Println(mergeSortedArr(a, b))
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
			if arr1Cur + 1 == len(arr1) {
				break
			}
			arr1Cur++
		} else {
			tmpArr = append(tmpArr, arr2[arr2Cur])
			if arr2Cur + 1 == len(arr2) {
				break
			}
			arr2Cur++
		}


	}

	return tmpArr
}
