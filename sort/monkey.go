package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := []int{1,5,3,6,8,1,5}
	for !isOrderd(a) {
		a = monkey(a)
	}

	fmt.Println(a)

}


func isOrderd(a []int) bool{
	pre := a[0]
	for _,v := range a{
		if v < pre {
			return false
		}
		pre = v

	}

	return true
}

func monkey(array []int) []int{
	exchangeCounts := 4 * len(array)
	for i := 0; i < exchangeCounts; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		from := int(r.Intn(len(array)))
		to := int(r.Intn(len(array)))
		var tmp int
		tmp = array[from]
		array[from] = array[to]
		array[to] = tmp
	}

	return array
}
