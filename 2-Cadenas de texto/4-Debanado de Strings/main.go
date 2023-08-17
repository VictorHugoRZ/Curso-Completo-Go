package main

import "fmt"

//Substrings o Debanado de Cadenas

func main() {
	var cadena string = "Este es un ejemplo de cadena de texto con debanado"

	fmt.Println(cadena[0]) //69 acorde al codigo assci

	fmt.Println(cadena[0:15]) //Este es un ejem
	//Imprime en la consola desde el primer elemento hasta el quinceavo, no acorde a las posiciones en el string.

	fmt.Println(cadena[:20]) //Este es un ejemplo d
	//Podemos omitir el primer valor y se tomara desde el primer elemento hasta donde nosotros indiquemos.
	//Al igual que en el ejemplo anterior, toma hasta el veinteavo elemento del string, no hasta la veinteava posicion del string.

	fmt.Println(cadena[20:]) //e cadena de texto con debanado
	//Podemos solo agregar el primer valor y omitir el segundo, esto tomara los todos los elementos que esten despues del veinteavo elemento de string.
	//Recordemos que tomara todos los elementos que resten despues del numero del elemento que le indiquemos.
}
