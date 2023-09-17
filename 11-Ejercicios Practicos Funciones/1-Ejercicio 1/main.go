package main

import "fmt"

func suma(numero1, numero2 int) int {
	resultadoSuma := numero1 + numero2
	return resultadoSuma
}

func resta(numero1, numero2 int) int {
	resultadoResta := numero1 + numero2
	return resultadoResta
}

func multiplicacion(numero1, numero2 int) int {
	resultadoMultiplicacion := numero1 * numero2
	return resultadoMultiplicacion
}

func division(numero1, numero2 int) int {
	resultadoDivision := numero1 / numero2
	return resultadoDivision
}

func main() {
	var operacion string

	fmt.Println("Calculadora")
	fmt.Println("¿Que operacion desea realizar?")
	fmt.Println("Ejemplo: suma, resta, multiplicacion, division")
	fmt.Scanln(&operacion)

	switch operacion {
	case "suma":
		fmt.Println("El resultado de la suma es:", suma(10, 10))
	case "resta":
		fmt.Println("El resultado de la resta es:", resta(100, 50))
	case "multiplicacion":
		fmt.Println("El resultado de la multiplicacion es:", multiplicacion(8, 8))
	case "division":
		fmt.Println("El resultado de la division es:", division(81, 9))
	default:
		fmt.Println("Ingresa una operacion valida")
	}
}
