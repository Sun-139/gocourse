package basics

import "fmt"

func main() {

	// Arithmatic operators

	//Addition Operator
	var a, b int = 3, 7
	var result = a + b
	fmt.Println("Addition : ", result)

	//Substraction operator
	result = a - b
	fmt.Println("Substraction : ", result)

	//Multiplication operator
	result = a * b
	fmt.Println("Multiplication : ", result)

	//Division operator
	result = a / b
	fmt.Println("Divion : ", result)

	//Modulus operator
	result = a % b
	fmt.Println("Remainder : ", result)

	const L = 3 / 7.0

	fmt.Println(L)

	//Overflow

	var x uint8 = 255
	fmt.Println("Before :", x)

	x = x + 1

	fmt.Println("After :", x)

	//Underflow

	var y uint8 = 0
	fmt.Println("Before :", y)

	y = y - 1
	fmt.Println("After :", y)

}
