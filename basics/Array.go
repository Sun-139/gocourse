package basics

import "fmt"

func main() {
	// num := [5]int{1, 2, 3, 4, 5}

	// for i := 4; i >= 0; i-- {
	// 	fmt.Println(num[i])
	// }

	// 2d array

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if j == 0 {
				continue
			}
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println(" ")
	}

}
