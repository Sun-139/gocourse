package basic

import "fmt"

func main() {

	//Printing numbers from 1 to 9 using loops
	fmt.Println("Numbers from 1 to 9")
	for i := 1; i <= 9; i++ {
		fmt.Println(i)
	}
	fmt.Println("")

	//Printing even numbers
	fmt.Println("Even numbers")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("")

	//Printing odd numbers
	fmt.Println("Odd numbers")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("")

	//Printing table of 9
	fmt.Println("Table of 9")
	for i := 1; i <= 10; i++ {
		fmt.Println(i, "* 9", "=", i*9)
	}
}
