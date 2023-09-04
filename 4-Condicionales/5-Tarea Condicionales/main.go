package main

import "fmt"

// Francisco quiere una taza de café, entonces la toma una taza y prepara su café,
// el solo toma su café cuando son las 3:00 pm de la tarde de lo contrario pasadas
// estas horas no toma café.
// ¿Cómo representarías este problema en el código?

func main() {
	var hora int8

	fmt.Println("¿Que hora es?")
	fmt.Scanln(&hora)

	if hora == 3 || hora == 03 {
		fmt.Println("Hora de tomar cafe!")
	} else {
		fmt.Println("Aun no es hora de tomar cafe")
	}
}
