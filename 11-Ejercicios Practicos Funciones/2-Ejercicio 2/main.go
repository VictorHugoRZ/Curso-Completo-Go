package main

import (
	"fmt"
	"strconv"
)

/*
Esta es una manera de transformar el tipo de dato entero a string para poder retornar los valores unidos
y asi Go no nos lance un error al unirlos
*/
func Persona(nombre string, edad int) (string, string) {
	stringEdad := strconv.Itoa(edad) //Convertimos el entero en string con este metodo de la libreria strconv
	return "Mi nombre es", nombre + " y tengo " + stringEdad + " años."
}

/*
Esta es otra manera mas sencilla de delcarar que tipo de valores retornara una funcion
*/
func Persona1(nombre string, edad int8) (string, string, string, int8, string) {
	return "Hola", nombre, ",tienes", edad, "años."
}

func main() {
	fmt.Println(Persona("Victor", 25))  //Mi nombre es Victor y tengo 25 años.
	fmt.Println(Persona1("Victor", 25)) //Hola, soy Victor y tengo 25 años.
	/*------------------------------------------------------------------------*/

	/*Aqui creamos variables para almacenar los datos que usaremos cuando invocamos la funcion*/
	var nombre string
	var edad int8
	fmt.Println("Bienvenido")
	fmt.Println("Por favor ingresa tu nombre:")
	fmt.Scan(&nombre)
	fmt.Println("Ingresa tu edad:")
	fmt.Scan(&edad)
	fmt.Println(Persona1(nombre, edad))

}
