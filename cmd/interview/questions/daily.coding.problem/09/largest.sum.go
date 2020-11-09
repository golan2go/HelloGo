package main

import "math"

// Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.
//For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.



func main() {
	arr := []int{5, 1, 1, 5}
	//nums := make([]int, len(arr))

	sum2 := arr[0]
	sum1 := int(math.Max(float64(arr[0]), float64(arr[1])))
	//nums = append(nums, sum1)
	for i:=2 ; i<len(arr) ; i++ {
		if sum2+arr[i] > sum1 {
			sum2, sum1 = sum1, sum2+arr[i]
		} else {
			sum2, sum1 = sum1, sum2+arr[i]
		}
	}

}
