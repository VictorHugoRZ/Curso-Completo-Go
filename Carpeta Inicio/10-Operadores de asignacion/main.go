package main

import "fmt"

func main() {
	numero := 20        //Se crea y se asigna la variable
	fmt.Println(numero) //20

	numero = 10         //Se reasigna la variable
	fmt.Println(numero) //10

	numero += 5         //Reasigna y suma 5 al valor total de numero
	fmt.Println(numero) //15

	numero -= 5         //Se reasigna y restan 5 al valor total del numero
	fmt.Println(numero) //10

	numero *= 5         //Se reasigna y multiplica el valor total del numero por 5
	fmt.Println(numero) //50

	numero /= 2         //Se reasigna y divide el valor total de numero en dos
	fmt.Println(numero) //25
}
