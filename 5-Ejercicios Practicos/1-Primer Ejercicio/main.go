package main

import "fmt"

func main() {
	var nombre string
	var transporte string

	fmt.Println("Vamos de viaje!")
	fmt.Println("¿Cual es tu nombre?")
	fmt.Scanln(&nombre)
	fmt.Println("-")
	fmt.Println("Bienvenido " + nombre + "!")
	fmt.Println("Nuestro destino es México, ¿Por que medio quieres llegar?")
	fmt.Println("¿Por Avion o Barco?")
	fmt.Scanln(&transporte)
	fmt.Println("-")
	fmt.Println("Excelente eleccion!")

	if transporte == "Avion" || transporte == "avion" {
		fmt.Println("El tiempo estimado para llegar a tu destino es de 72 horas, disfruta tu vuelo!")
	} else if transporte == "Barco" || transporte == "barco" {
		fmt.Println("El tiempo estimado para llegar a tu destino es de 120 horas, disfruta del mar!")
	} else {
		fmt.Println("Por favor elige una via de transporte valida")
	}
}
