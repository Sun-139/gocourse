package basics

import "fmt"

func main() {

	fmt.Println("2D Arrays with range loop")

	//Print all values using matrix style
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for _, row := range matrix {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
	fmt.Println()

	//Print only second Row
	for i, row := range matrix {
		if i == 0 {
			continue
		}
		fmt.Println(row)
	}
	fmt.Println()

	//Print only first Column

	for _, row := range matrix {
		for j, val := range row {
			if j != 0 {
				continue
			}
			fmt.Println(val)
		}
	}
	fmt.Println()

	//Find sum of all element

	var sum int
	for _, row := range matrix {
		for _, val := range row {
			sum = val + sum
		}
	}
	fmt.Println("The sum of element is", sum)
	fmt.Println()

	//Find largest number in 2D array

	lar := 0
	for _, row := range matrix {
		for _, val := range row {
			if val > lar {
				lar = val
			}
		}
	}
	fmt.Println("The largest element is", lar)

}
