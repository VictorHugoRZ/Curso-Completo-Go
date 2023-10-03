package main

import "fmt"

/*2- En esta funcion declaramos que se recibiran dos parametros, el primero sera un String, el segundo sera
un canal junto con su tipo de dato compatible de acuerdo a los parametros anteriores, al nosotros agregar
 <- estamos declarando que el canal que ingresaremos por la funcion recibira informacion y al declarar la
sentencia de codigo y querer introducir informacion al canal 1 de nuestro canal c, agregaremos el mismo
signo <- para declarar que la variable msg tendra una entrada al canal c*/
func mensaje(msg string, c chan<- string) {
	c <- msg
}

func main() {

	/*1- Aqui definimos nuestro canal de informacion con el metodo Make(), ademas declaramos que el tipo de
	dato que tendra nuestro canal es de tipo String y el numero define el numero de canales internos dentro
	nuestro canal*/
	c := make(chan string, 1)

	/*3- Invocamos a nuestra funcion de la linea 10 para introducir nuestro valor al canal que creamos en la
	linea 19*/
	go mensaje("Estoy en el canal", c)

	/*4- Al declarar nosotros <-c estamos definiendo el valor que se imprimira con el metodo fmt, este sera el
	valor de salida del canal c*/
	fmt.Println(<-c)
}
