package main

import "fmt"

func main() {
	var nombre string
	var direccion string
	var edad int8
	var profesion string

	fmt.Println("Introduce tu nombre: " + nombre)
	fmt.Scanln(&nombre)

	fmt.Println("Introduce tu direccion: " + direccion)
	fmt.Scanln(&direccion)

	fmt.Println("Introduce tu edad: ")
	fmt.Scanln(&edad)

	fmt.Println("Introduce tu profesion: " + profesion)
	fmt.Scanln(&profesion)
}
