package main

import (
	"fmt"
	"reflect"
)

func main() {
	var cifra int = 10
	var palabra string = "Victor"
	fmt.Println(cifra)
	fmt.Println(palabra)
	fmt.Println(reflect.TypeOf(cifra))
	fmt.Println(reflect.TypeOf(palabra))
}
