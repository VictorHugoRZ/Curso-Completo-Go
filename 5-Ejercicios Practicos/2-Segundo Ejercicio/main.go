package main

import "fmt"

func main() {
	var juego string
	fmt.Println("Tenemos que decidir que tipo de juego crear...")
	fmt.Println("Existen dos opciones, elige una de ellas:")
	fmt.Println("-MMO")
	fmt.Println("-Shooter")
	fmt.Println("Elige con cuidado, tu eleccion determinara el trabajo de la empresa durante los proximos años")
	fmt.Println("¿Cual es tu eleccion?")
	fmt.Scanln(&juego)

	if juego == "MMO" || juego == "mmo" {
		fmt.Println("Perfecto!, tenemos grandes posibilidades de hacer un buen juego, manos a la obra!")
	} else if juego == "Shooter" || juego == "shooter" {
		fmt.Println("Buena eleccion!, comencemos a trabajar!")
	} else {
		fmt.Println("Por favor, elige el tipo de juego de nuevo")
	}
}
