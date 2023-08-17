package main

import "fmt"

func main() {
	var cadena string = "Este es un ejemplo de \"cadena de texto\""
	//Si queremos agregar comillas dobles a nuestros strings en go debemos agregar barras invertidas antes de esas comillas.
	fmt.Println(cadena)

	var cadena2 string = "Este es un \nsalto de linea"
	//Para agregar saltos de linea en nuestros strings podemos agregar una barra invertida junto con una "n", es decir: \n.
	fmt.Println(cadena2)

	var cadena3 string = "Este \tes \tun \tejemplo \tde \ttabulacion"
	//La barra invetida mas la letra "t" genera espacios equivalentes a una tabulacion en nuestros strings \t.
	fmt.Println(cadena3)

	var cadena4 string = "Este es el ejemplo para borrar espacios en nuestros \bstrings"
	//La barra invetida mas la letra "b" elimina los espacios que existen entre nuestros strings \b.
	fmt.Println(cadena4)

	var cadena5 string = "Este es el ejemplo para colocar los simbolos de los generos \vmasculino y \ffemenino"
	//Con estos signos agregamos los simbolos de masculino y femenino a nuestros strings.
	fmt.Println(cadena5)

	var cadena6 string = "Este es el ejemplo de texto \rrandom"
	//Altera el string en modo aleatorio.
	fmt.Println(cadena6)

}
