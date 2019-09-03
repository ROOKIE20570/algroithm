package main

import "fmt"

func main() {
	a := []int{1, 5, 2, 6, 8, 4,}
	fmt.Println(selection(a))
}

func bubble(a []int) []int {
	for i := 0; i < len(a); i++ {
		flag := false
		for j := 0; j < len(a)-i-1; j++ { //每次必然会把最大的摆到最后 所以执行到这的时候 有i+1位已经被排好序了
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}

		if !flag {
			return a
		}
	}

	return a
}

func selection(a []int) []int {

	for i := 0; i < len(a); i++ {
		minIndex := i
		for j := minIndex; j < len(a); j++ {
			if a[minIndex] > a[j] {
				minIndex = j
			}
		}

		a[i], a[minIndex] = a[minIndex], a[i]
	}

	return a

}

//插入排序
func insert(a []int) []int {
	b := 1
	for ; b < len(a); b++ {
		for i := 0; i < b; i++ {
			if a[b] < a[i] {
				a[b], a[i] = a[i], a[b]
			}
		}
	}

	return a
}
