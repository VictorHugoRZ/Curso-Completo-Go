package main

import "fmt"

/*Podemos definir los parametros para nuestras funciones declarandolos dentro de los parentesis,
tambien debemos definir el tipo de dato de cada uno de esos parametros*/

func suma(numero int8) {
	fmt.Println(numero * 2) //20
}

func mensajeUsuario(mensaje string) {
	fmt.Println(mensaje) //Hola, mi nombre es Victor
}

//En la funcion saludo usamos una variable
var nombre string = "Victor"

func saludo() {
	fmt.Println(nombre) //Victor
}

//---------------------------------------------------------------------------------------------------//

func retorno() string {
	saludo := "Hola Victor"
	return saludo
}

func formulario(nombre string, apellido string) string {
	return "Hola!, mi nombre es " + nombre + " " + apellido
}

func main() {
	suma(10)                                     //20
	mensajeUsuario("Hola, mi nombre es Victor")  //Hola, mi nombre es Victor
	saludo()                                     //Victor
	var llamada string = retorno()               //Llamamos a la funcion en esta variable y despues imprimimos el resultado
	fmt.Println(llamada)                         //Hola Victor
	fmt.Println(formulario("Victor", "Ramirez")) //Hola!, mi nombre es Victor Ramirez
}
