package main

import "fmt"

func main() {
	//Los arrays en Go son variables a las que les indicaremos el limite de espacio que tendran
	//Tenemos que indicar dentro de las llaves el numero de espacios que este array tendra, de lo contrario
	//no podremos almacenas nada dentro del mismo y todos los datos que querramos almacenar deben ser
	//de el mismo tipo de dato

	var personas [5]string
	fmt.Println(personas) //[]

	personas[0] = "Alejandro"
	fmt.Println(personas) //[Alejandro    ]

	personas[1] = "Leonardo"
	fmt.Println(personas) //[Alejandro Leonardo   ]

	personas[2] = "Patricia"
	personas[3] = "Paola"
	personas[4] = "Amparo"
	fmt.Println(personas) //[Alejandro Leonardo Patricia Paola Amparo]

	//Los espacios que quedan despues de llamar a imprimir la informacion por consola en las lineas 15 y 18
	//Son el numero de espacios que aun tenemos libres dentro de nuestro array para insertar informacion
	//En este caso nuestro array puede contener 5 elementos de tipo string y al completar estos 5 elementos
	//dentro del array en la linea 22, al llamar a imprimir la informacion por consola ya no tenemos espacios
	//libres para insertar mas datos

	fmt.Println(personas[4]) //Amparo
	//De esta manera podemos consultar informacion dentro de nuestro array

	familia := [2]string{"Sabrina", "Victor"}
	fmt.Println(familia) //[Sabrina Victor]
	//De esta manera podemos declarar arrays dandoles un valor inicial

	familia[0] = "Alejandra"
	fmt.Println(familia) //[Alejandra Victor]
	//De esta manera reasignamos valores dentro del array

}
