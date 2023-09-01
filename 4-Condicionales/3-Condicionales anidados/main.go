package main

import "fmt"

func main() {
	var nombre string
	var edad int8

	fmt.Println("Ingresa tu nombre:")
	fmt.Scanln(&nombre)

	if nombre == "Victor" {
		fmt.Println("Ingresa tu edad:")
		fmt.Scanln(&edad)

		if edad >= 18 {
			fmt.Println("Bienvenido!")
		} else {
			fmt.Println("Lo siento, no tienes la mayoria de edad")
		}

	} else {
		fmt.Println("Lo siento, no estas en la lista de invitados")
	}

}
