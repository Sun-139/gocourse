package basics

import "fmt"

func main() {

	fmt.Println("Practicing on slices")

	//Printing all elements of slice using range
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range nums {
		fmt.Println(v)
	}
	fmt.Println("")

	//Find sum using range

	var sum int
	for _, v := range nums {
		sum += v

	}
	fmt.Println("The sum is", sum)
	fmt.Println()

	//Craet new slice with only numbes > 3

	var greater []int
	for _, v := range nums {
		if v > 3 {
			greater = append(greater, v)
		}
	}
	fmt.Println(greater)
	fmt.Println()

	//Multiply all elements by 3 { modify the original slice}

	for i := range nums {
		nums[i] = nums[i] * 3
	}
	fmt.Println("This is the modify slice")
	fmt.Println(nums)
	fmt.Println()

	//Remove element at index 1 using slicing

	index := 2
	nums = append(nums[:index], nums[index+1:]...)
	fmt.Println(nums)
}
