package main

import "fmt"

func main() {
	//Break
	//La palabra clave break nos permite detener la ejecucion de nuestro codigo en un punto especifico

	for paquete := 1; paquete <= 20; paquete++ {
		if paquete == 15 {
			fmt.Println("Hemos tenido un inconveniente con el paquete", paquete, "reanudaremos las entregas el dia de mañana.")
			break
		}
		fmt.Println("El paquete", paquete, "ha sido entregado.")
	}
}

//Gracias a la keyword break podemos cortar la ejecucion de nuestro codigo, en este caso con este bucle for
//que nos permite detenernos cuando la variable paquete llega a tener un valor de 15
