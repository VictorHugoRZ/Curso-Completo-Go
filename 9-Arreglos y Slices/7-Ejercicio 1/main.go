package main

import "fmt"

func main() {
	var articulo string
	carrito := []string{}

	for i := 0; i < 10; i++ {
		if i == 0 {
			fmt.Println("¿Desea agregar algún artículo a su carrito?")
			fmt.Scanln(&articulo)
			carrito = append(carrito, articulo)
		} else if i > 0 {
			fmt.Println("¿Desea agregar algo más?")
			fmt.Scanln(&articulo)
			carrito = append(carrito, articulo)
		}
	}

	fmt.Println("Su carrito está lleno, los articulos que lleva son", carrito)
	fmt.Println("Gracias por su compra!")
}
