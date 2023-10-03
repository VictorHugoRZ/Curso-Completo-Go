package main

import "fmt"

func main() {

	/*Declaramos nuestro canal definiendo su tipo de dato como String, ademas definimos que tendra
	3 sub canales dentro de si*/
	c := make(chan string, 3)

	c <- "Hola Sabrina"   //Arbitrariamente se asigna el primer valor
	c <- "Hola Leonardo"  //Arbitrariamente se asigna el segundo valor
	c <- "Hola Alejandro" //Arbitrariamente se asigna el tercer valor

	/*Aqui usamos el metodo len() para saber la longitud de sub canales que se encuentran ocupados de
	nuestro canal c, ademas usamos el metodo cap() para saber el total de sub canales que nuestro canal c
	tiene*/
	fmt.Println(len(c), cap(c)) // 3 3

	close(c) /*Con el metodo close() podemos cerrar la entrada de datos o valores de cualquier canal que
	tengamos en nuestro codigo, esto hara que el canal ya no admita datos nuevos o sobre escritos*/

	/*Imprimimos con un for range cada valor de los sub canales de nuestro canal c*/
	for valores := range c {
		fmt.Println(valores)
	}

}
