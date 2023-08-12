package main

import "fmt"

func main() {
	var num1 int = 20
	var num2 int = 10

	fmt.Println(num1 + num2) //30
	fmt.Println(num1 - num2) //10
	fmt.Println(num1 * num2) //200
	fmt.Println(num1 / num2) //2
	fmt.Println(num1 % num2) //0

	num2++            //10 + 1
	fmt.Println(num2) //11

	num1--            //20 - 1
	fmt.Println(num1) //19

	fmt.Println("Este es un valor negativo con el operador unario: ", -num1)
	//Al conocar el signo de más + o menos - cuando llamamos a una funcion, esto convertira el valor de
	//esa variable en negativo o positivo, a excepcion de que el signo ya sea negativo, ahi se
	//aplicara la ley de signos y sera positivo.
}
