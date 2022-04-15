package main

import (
	"fmt"
)

func main() {
	//arr := []int{-5, -10, 0, -3, 8, 5, -1, 10}
	arr := []int{1,8,9,0 ,0,1,1,1,2,10,-1}

	var min, max = 100000000, -1

	for _, x := range arr {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	fmt.Println(max)
	fmt.Println(min)
	count := make([]int, max-min+1)
	fmt.Println(count)

	for i := 0; i < len(arr); i++ {
		count[arr[i]-min]++
	}

	fmt.Println(count)
	for i := 1; i <= len(count)-1; i++ {
		count[i] = count[i] + count[i-1]
	}

	fmt.Println(count)
	arr2 := make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		arr2[count[arr[i]-min]-1] = arr[i]
		count[arr[i]-min]--
	}

	  fmt.Println(arr2)
}
