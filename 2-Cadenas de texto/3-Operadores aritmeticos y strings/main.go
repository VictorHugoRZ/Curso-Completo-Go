package main

import "fmt"

func main() {
	cadena1 := "Hola Victor"
	cadena2 := "Ramirez"
	var numero int = 1

	fmt.Println(cadena1 + cadena2) //Hola VictorRamirez

	fmt.Println("Hola Usuario: ", cadena2) //Hola Usuario:  Ramirez

	fmt.Println("Hola Usuario: " + cadena2) //Hola Usuario: Ramirez

	fmt.Println(numero, "Usuario: "+cadena2) //1 Usuario: Ramirez
}
