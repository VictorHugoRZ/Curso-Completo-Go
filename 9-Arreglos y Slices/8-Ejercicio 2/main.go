package main

import "fmt"

func main() {
	var decision string
	var producto string
	carrito := []string{
		"Pepitas",
		"Pan de caja",
		"Ajonjolí",
	}

	fmt.Println("¿Desea agregar algo mas a su carrito?")
	fmt.Scanln(&decision)

	if decision == "no" || decision == "No" || decision == "NO" {
		fmt.Println("Perfecto, estos son los productos de su carrito hasta ahora:", carrito)
	} else if decision == "si" || decision == "Si" || decision == "SI" {
		for i := 1; i <= 5; i++ {
			fmt.Println("Esta es su lista hasta el momento:", carrito)
			fmt.Println("¿Que más desea agregar?")
			fmt.Scanln(&producto)
			carrito = append(carrito, producto)
		}

		fmt.Println("\nTodos los productos han sido agregados correctamente, aquí esta su lista:", carrito, "\nGracias por su compra!")
	}
}
