package basics

import "fmt"

func main() {

	fmt.Println("this is Practice of arrays and slices")

	// num := [9]int{1, 2, 3, 4, 5, 6, 7, 9}
	// fmt.Println(num[3])

	arr2 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i := 0; i < 2; i++ {
		// if i == 0 {
		// 	continue
		// }
		for j := 0; j < 3; j++ {
			if j == 0 {
				continue
			}
			fmt.Println(arr2[i][j])
		}
		// fmt.Println(arr2[i][1])
	}
	fmt.Println("")
	fmt.Println("Printing slices in reverse")
	num1 := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 6; i >= 0; i-- {
		fmt.Println(num1[i])
	}
	fmt.Println("")
	fmt.Println("largest number finding")

	num2 := []int{1, 20, 3, 44, 52, 64, 76}
	var lar int
	for i := 0; i < len(num2); i++ {
		if num2[i] > lar {
			lar = num2[i]
		}
	}
	fmt.Println(lar)

	fmt.Println("")
	fmt.Println("Print even number from slice")

	num3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(num3); i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(num3[i])
	}
	fmt.Println("")
	fmt.Println("Print odd number from slice")

	num4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(num4); i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(num4[i])
	}
	fmt.Println("")
	fmt.Println("Creating a Slice from Slice")
	fruitList := []string{"Mango", "Apple", "Guava", "Banana"}
	var fruitList2 []string
	for i := 0; i < len(fruitList); i++ {
		fruitList2 = append(fruitList2, fruitList[i])
	}
	fmt.Println(fruitList2)

}
