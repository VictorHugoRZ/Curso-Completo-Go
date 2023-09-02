package main

import "fmt"

func main() {
	var numero int8

	// switch <variable a evaluar> {
	// case <valor>:
	// 	<sentencia>
	// case <valor>:
	// 	<sentencia>
	// default:
	// 	<sentencia>
	// }

	fmt.Println("Introduce un numero:")
	fmt.Scanln(&numero)

	switch numero {
	case 1:
		fmt.Println("Este es el primer valor de la serie")
	case 2:
		fmt.Println("Este es el segundo valor de la serie")
	case 3:
		fmt.Println("Este es el tercer valor de la serie")
	case 4:
		fmt.Println("Este es el cuarto valor de la serie")
	case 5:
		fmt.Println("Este es el quinto valor de la serie")
	default:
		fmt.Println("Este valor esta fuera del rango de esta serie numerica")
	}
}
