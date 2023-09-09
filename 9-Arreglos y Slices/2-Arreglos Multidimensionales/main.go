package main

import "fmt"

func main() {
	numeros1 := [5]int8{1, 2, 3, 4, 5}
	numeros2 := [5]int8{6, 7, 8, 9, 10}
	numeros3 := [5]int8{11, 12, 13, 14, 15}

	multidimensional := [][5]int8{numeros1, numeros2, numeros3}
	fmt.Println(multidimensional) //[[1 2 3 4 5] [6 7 8 9 10] [11 12 13 14 15]]
	//Podemos observar que ahora tenemos distintos arrays dentro de otro array que los almacena todos

	/*Al declarar arrays multidimensionales tenemos que incluir dos corchetes, el primero sera para
	declarar el numero de espacios que tendra este array multidimensional, el segundo para declarar
	el numero de espacios que deben tener los arrays que seran incluidos, todo esto para ser
	totalmente precisos en cuando a el espacio que ocuparemos
	Aunque es recomendable que no limitemos el espacio de elementos del array multidimensional
	Ya que podremos utilizarlos para almacenar grandes cantidades de informacion*/

	nombres := [2]string{"Sabrina", "Victor"}
	apellidos := [2]string{"Torres", "Ramirez"}

	multi2 := [3][2]string{nombres, apellidos}
	fmt.Println(multi2) //[[Sabrina Victor] [Torres Ramirez] [ ]]
	/*En este ultimo array multidimensional tenemos un espacio en blanco al final del mismo, esto se debe
	a que cuando declaramos el array multidimensional en la linea 24, tambien declaramos que el array
	tendria un total de 3 elementos y que cada uno de esos elementos debe tener dos elementos dentro de si*/

	fmt.Println(multi2[0][0]) //Sabrina
	/*De esta manera podemos acceder a valores dentro de arrays multidimensionales, primero indicamos el
	numero del elemento al que queremos acceder dentro del array multidimensional y despues indicamos el
	numero del elemento al que queremos acceder dentro de este sub array*/

}
