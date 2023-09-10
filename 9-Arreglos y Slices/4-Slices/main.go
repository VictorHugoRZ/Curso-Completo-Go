package main

import "fmt"

func main() {
	/*Los Slices son arreglos que no tienen definido el numero de elementos que se almacenaran en ellos,
	esto nos permite usarlos para guardar grandes cantidades de informacion*/

	slice := []string{"Victor", "Hugo", "Ramirez", "Zamora"}

	fmt.Println(slice)      //[Victor Hugo Ramirez Zamora]
	fmt.Println(slice[0:2]) //[Victor Hugo]
	/*Esto nos permite consultar valores dentro de los arrays y slices, el primer numero nos indica
	desde que posicion tomara el primer valor, el segundo numero nos indica que hasta esa posicion
	tomara el valor anterior, por eso toma hasta Hugo en la linea 12*/

	fmt.Println(slice[1:3]) //[Hugo Ramirez]
	//El primer valor es inclusivo y el segundo es exclusivo
}
