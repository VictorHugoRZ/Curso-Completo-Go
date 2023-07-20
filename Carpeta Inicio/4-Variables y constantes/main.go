package main

import "fmt"

func main() {
	var variable int = 1 // Estructura similar a Typescript < var "nombre" "tipo de dato" = "valor" >
	fmt.Println(variable)

	variable = 25 // Reasignacion de valor
	fmt.Println(variable)

	variable = 100 // Reasignacion de valor
	fmt.Println(variable)

	const constante = 9.8  // Estructura diferente a variables normales < const "nombre" = "valor" >
	fmt.Println(constante) // No podemos reasignar valores a constantes
}
