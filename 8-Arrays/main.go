package main

import "fmt"

func main() {

	//Arrays
	// var nums [5]int
	// nums[0] = 1
	// fmt.Println(nums[0])
	// fmt.Println(nums)

	nums := [5]int{1, 2, 3}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	//2dArrays
	nums2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(nums2)

}
