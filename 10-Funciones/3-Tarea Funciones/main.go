package main

import "fmt"

/*Juan necesita saber cómo implementar dentro de su código una solución para determinar
el área de un triángulo, pero este debe de hacer que el usuario pueda ingresar diferentes
valores para poder realizar el cálculo y también no puede permitirse el repetir variables
y datos cada vez que se ingresen.*/

var base, altura int8

func area(base int8, altura int8) int8 {
	return base * altura / 2
}

func main() {
	fmt.Println("Calculadora de área de triángulos")
	fmt.Println("Introduce los valores del triángulo")
	fmt.Println("Base:")
	fmt.Scanln(&base)
	fmt.Println("Altura:")
	fmt.Scanln(&altura)
	var resultado int8 = area(base, altura)
	fmt.Println("El resultado es:")
	fmt.Println(resultado)
}
