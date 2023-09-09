package main

import "fmt"

func main() {
	//Defer
	//La palabra clave defer nos permite "saltar" una seccion de codigo determinada para que esta se ejecute
	//despues de todo el codigo siguiente.

	defer fmt.Println("Este mensaje va primero")
	fmt.Println("Este mensaje va segundo")
	fmt.Println("Este mensaje va tercero")
	fmt.Println("Este mensaje va cuarto")

	// Este mensaje va segundo
	// Este mensaje va tercero
	// Este mensaje va cuarto
	// Este mensaje va primero

	//No se recomienda utilizar mas de un defer en nuestro codigo, ya que el efecto de este se puede volver
	//un poco impredecible
}
