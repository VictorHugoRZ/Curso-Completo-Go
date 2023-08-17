package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hola")

	var cadena string = "Hola Victor"

	fmt.Println(reflect.TypeOf(cadena))
}
