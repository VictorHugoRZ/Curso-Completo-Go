package main

import "fmt"

func main() {
	var edad int8 = 18

	// if <condicion o condiciones> {
	// 	<sentencia de codigo>
	// } else {
	// 	<sentencia de codigo>
	// }

	// Primera manera de hacer condicionales

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Printf("No eres mayor de edad")
	}

	// Segunda manera de hacer condicionales
	// En esta segunda manera podemos declarar las variables a evaluar dentro del mismo If

	if edad := 68; edad > 60 {
		fmt.Println("Tienes mas de sesenta años")
	} else {
		fmt.Println("No tienes la edad requerida")
	}
}
