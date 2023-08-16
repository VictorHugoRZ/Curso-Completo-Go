package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hola")

	var cadena string = "Hola Mundo"

	fmt.Println(reflect.TypeOf(cadena))
}
