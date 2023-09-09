package main

import "fmt"

func main() {
	array := [3]int8{1, 2, 3}
	fmt.Println(array) //[1 2 3]

	/*Para conocer la longitud total de un arreglo podemos utilizar el metodo len, recordemos que esto nos
	dara el numero total de elementos que el array podra almacenar, aunque alguno de sus espacios aun no este
	ocupado, NO nos dara el numero de elementos actuales en el array, este ultimo valor puede variar
	y podria ser inexacto*/

	fmt.Println(len(array)) //3

	array2 := [5]int8{1, 2, 3, 4}
	fmt.Println(array2)      //[1 2 3 4 0]
	fmt.Println(len(array2)) //5
	fmt.Println(cap(array2)) //5
	//Cap nos permite saber la capacidad del array
}
