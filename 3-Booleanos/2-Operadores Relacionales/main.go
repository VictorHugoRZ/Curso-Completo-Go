package main

import "fmt"

func main() {
	num1 := 20
	num2 := 30
	fmt.Println(num1 < num2) //true
	fmt.Println(num1 > num2) //false

	num1 = 30
	fmt.Println(num1 <= num2) //true

	num1 = 31
	fmt.Println(num1 >= num2) //true
	fmt.Println(num1 == num2) //false
	fmt.Println(num1 != num2) //true
}
