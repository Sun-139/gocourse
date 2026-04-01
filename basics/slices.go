package basics

import (
	"fmt"
	"slices"
)

func main() {

	//slices declaration
	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Println("its a slice", slice)

	//printing slice
	nums := []int{1, 2, 3, 4, 5}
	for i, v := range nums {
		fmt.Println(i, v)
	}
	fmt.Println(" ")

	//print in reverse
	for i := 4; i >= 0; i-- {
		fmt.Println(nums[i])
	}
	fmt.Println(" ")

	//sum of elements
	var sum int
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	fmt.Println("the sum of elements is ", sum)

	//Printing Largest Number
	num2 := []int{11, 45, 22, 33, 78, 63, 24}
	fmt.Println(slices.Max(num2))

	fmt.Println(" ")

	//Printing Even numbers
	num3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for i := 0; i < len(num3); i++ {
		if num3[i]%2 != 0 {
			continue
		}
		fmt.Println(num3[i])
	}
	fmt.Println(" ")

	//creat new slice with odd numbers
	var num4 []int
	for i := 0; i < len(num3); i++ {
		if num3[i]%2 == 0 {
			continue
		}
		num4 = append(num4, num3[i])

	}
	fmt.Println(num4)
}
