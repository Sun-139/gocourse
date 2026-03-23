package basics

import "fmt"

func main() {

	age := 25
	if age >= 18 {
		fmt.Println("person is adult")
	}

	num := 18
	if num%2 == 0 {
		if num%3 == 0 {
			fmt.Println("Number is divisible by both 2 and 3")
		} else {

			fmt.Println("Number is divisible by 2 but no 3")
		}

	} else {

		fmt.Println("Number is not divisible by both")
	}
	//problem solving
	count := 0
	for i := 1; i <= 30; i++ {
		if i%6 == 0 {
			continue
		}
		if i%3 == 0 {
			count++
		}
	}
	fmt.Println(count)

	num1 := 14
	if num1%2 == 0 && num1%7 == 0 {
		fmt.Println("Number is divisible by both numbers")
	}
	if num1%2 == 0 || num1%3 == 0 {
		fmt.Println("Either 2 or 3 is divisible")
	}
}
