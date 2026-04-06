package basics

import "fmt"

func main() {

	fmt.Println("Practice array with for range")
	fmt.Println("")

	//Print all values using range
	arr := [5]int{1, 2, 3, 4, 5}

	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println("")

	//Print only index values
	for i := range arr {
		fmt.Println(i)
	}
	fmt.Println("")

	//Find sum using range
	var sum int
	for _, v := range arr {
		sum = v + sum
	}
	fmt.Printf("The sum is %d", sum)
	fmt.Println("")
	fmt.Println("")

	//Convert even numbers using range
	arr2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range arr2 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(arr2[i])
	}
	fmt.Println("")

	//Multiply each element by 2
	for _, v := range arr2 {
		fmt.Println(v, "*", "2", "=", v*2)
	}
}
