package main

import "fmt"

func main() {
	var nombre string
	var edad int

	fmt.Println("Ingresa tu nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Ingresa tu edad: ")
	fmt.Scanln(&edad)

	fmt.Println("Tu nombre es", nombre, "y tienes", edad, "años")
}
