package main

import (
	"fmt"
	"reflect"
)

func main() {
	var booleano bool = true

	fmt.Println(booleano) //true
	//Las variables booleanas que no tienen valor asignado por default tendran el valor de false

	fmt.Println(reflect.TypeOf(booleano)) //bool
}
