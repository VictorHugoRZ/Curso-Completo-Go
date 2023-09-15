package main

import "fmt"

func saludoTodos() {
	fmt.Println("Hola!")
	fmt.Println("Buen dia!")
}

func despedidaTodos() {
	fmt.Println("Adios!")
	fmt.Println("Buena noche!")
}

func main() {
	saludoTodos()
	despedidaTodos()
}

/*Podemos definir funciones con la palabra reservada func, el nombre de nuestra funcion, parentesis
y llaves para definir el cuerpo de la funcion*/
