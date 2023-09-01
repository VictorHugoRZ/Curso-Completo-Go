package main

import "fmt"

func main() {
	num1 := 10
	num2 := 20
	num3 := 30
	num4 := 40

	fmt.Println(num1 < num2 && num3 == num4) //false

	fmt.Println(num1 != num2 || num3 > num4) //true
}
