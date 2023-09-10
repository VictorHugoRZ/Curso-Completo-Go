package main

import "fmt"

func main() {
	numeros := []int8{1, 2, 3, 4, 5}
	fmt.Println(numeros) //[1 2 3 4 5]

	numeros = append(numeros, 6)
	fmt.Println(numeros) //[1 2 3 4 5 6]

	numeros2 := []int8{7, 8, 9, 10}
	fmt.Println(numeros2) //[7 8 9 10]

	numeros = append(numeros, numeros2...)
	fmt.Println(numeros) //[1 2 3 4 5 6 7 8 9 10]
	/*Para concatenar arreglos o slices utilizamos el operador de propagacion o rest operator, esto nos permite unir de manera correcta
	un array o slice junto con otro, el metodo append agrega informacion al final del array o slice*/

	numeros = append(numeros, numeros2[3])
	fmt.Println(numeros) //[1 2 3 4 5 6 7 8 9 10 10]
	//Tambien podemos agregar solo un elemento del array o slice que necesitemos accediendo mediante su posicion al valor necesario
}
