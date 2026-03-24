package basics

import "fmt"

func main() {

	// Switch case statements

	// fruit := "apple"

	// switch fruit {
	// case "mango":
	// 	fmt.Println("its a mango")
	// case "apple":
	// 	fmt.Println("its a apple")
	// default:
	// 	fmt.Println("its not a fruit")
	// }

	// Problem solving

	var num int
	fmt.Println("Enter a number between 1 - 7")
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("what the hell are you entering")
	}

	//Problem solving part 2

	var marks int
	fmt.Println("Enter the marks")
	fmt.Scanln(&marks)

	switch {
	case marks >= 90:
		fmt.Println("Grade A")
	case marks >= 75 && marks <= 89:
		fmt.Println("Grade B")
	case marks >= 50 && marks <= 74:
		fmt.Println("Grade C")
	default:
		fmt.Println("FAIL")
	}

}
