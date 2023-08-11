package main

import "fmt"

func main() {
	var num3ro int = 3  //Podemos omitir el tipo de dato de la variable, aun asi se asignara automaticamente
	fmt.Println(num3ro) //Las variables o constantes no pueden llevar numeros al inicio del nombre

	var num_3ro int = 20 //Pueden llevar numeros al final y en medio, asi como guiones bajos, no otro tipo
	fmt.Println(num_3ro) //de guiones

	//Las variables no pueden tener nombre de metodos o palabras reservadas del lenguaje

	//Usa camelCase para nombrar variables descriptivas

	/*Para programas grandes asignaremos los nombres de las variables en Mayuscula, esto hara posible que
	al importar datos en otros archivos esas variables tengan visibilidad en ellos*/
}
