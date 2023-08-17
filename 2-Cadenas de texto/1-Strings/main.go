package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hola")

	var cadena string = "Hola Victor" //Se usan comillas dobles para declarar strings y comillas simples
	//para caracteres individuales

	fmt.Println(reflect.TypeOf(cadena))
}
