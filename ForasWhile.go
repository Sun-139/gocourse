package main

import "fmt"

func main() {
	//Break Statement
	// sum := 0

	// for {
	// 	sum += 10
	// 	fmt.Println(sum)
	// 	if sum >= 50 {
	// 		break

	// 	}

	// }

	//Continue statment
	// num := 1
	// for num <= 10 {
	// 	if num%2 == 0 {
	// 		num++
	// 		continue

	// 	}
	// 	fmt.Println("odd number :", num)
	// 	num++
	// }

	// i := 1
	// for i <= 10 {
	// 	if i == 5 {
	// 		i++
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	i++

	// }
	// fmt.Println("")
	// j := 1
	// for j <= 10 {
	// 	if j == 7 {
	// 		j++
	// 		break
	// 	}
	// 	fmt.Println(j)
	// 	j++
	// }
	// fmt.Println("")

	// s := 1
	// for s <= 20 {
	// 	if s%3 == 0 {
	// 		s++
	// 		continue
	// 	}
	// 	fmt.Println(s)
	// 	s++
	// }

	// u := 1
	// for u <= 10 {
	// 	if u%4 == 0 {
	// 		u++
	// 		break
	// 	}
	// 	fmt.Println(u)
	// 	u++
	// }
	// fmt.Println("")

	// n := 1
	// for n <= 15 {
	// 	if n == 11 {
	// 		break
	// 	}
	// 	if n%2 == 0 {
	// 		n++
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// 	n++

	// }
	// fmt.Println("")

	// var p int
	// fmt.Println("enter a number")
	// fmt.Scanln(&p)

	// r := 1
	// for r <= 10 {
	// 	fmt.Println(r, "*", p, "=", r*p)
	// 	r++
	// }

	// i := 10
	// for i >= 1 {
	// 	i--
	// 	fmt.Println(i)
	// }

	var count = 0
	for i := 1; i <= 50; i++ {
		if i%5 == 0 {
			count++
		}
	}
	fmt.Println(count)

	var count2 = 0
	f := 1
	for f <= 50 {
		if f%5 == 0 {
			count2++
		}
		f++
	}
	fmt.Println(count2)
}
