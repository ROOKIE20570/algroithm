package main

import "fmt"

func main() {
	fmt.Println(1 >> 1)
}

func test(a []int) {
	a[1] = 5

}
func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}

	mergeSort(arr, 0, arrLen-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tmpArr)
}

func quickSort(arr []int) {
	quickSortC(arr, 0, len(arr)-1)
}

func quickSortC(arr []int, i, j int) {
	if i >= j {
		return
	}

	p := partition(arr, i, j)
	quickSortC(arr, i, p-1)
	quickSortC(arr, p+1, j)
}

func partition(arr []int, p, r int) int {
	pivot := arr[r] //取最后一个作为分界
	i := p          //暂存第一个元素
	for j := p; j < r; j++ { //迭代一圈数据
		if arr[j] < pivot { //如果比分界点小
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[r] = arr[r], arr[i]

	return i
}
