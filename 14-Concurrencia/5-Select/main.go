package main

import "fmt"

//Esta funcion asigna el valor del parametro de tipo String a el canal que le pasemos como segundo parametro
func mensaje(msg string, c chan<- string) {
	c <- msg
}

func main() {

	//Definimos nuestros canales de tipo String ambos
	Canal1 := make(chan string)
	Canal2 := make(chan string)

	//Definimos nuestras Goroutines
	/*Una Goroutine es un proceso especifico al que definimos con la palabra reservada go.
	Estos se quedaran ejecutandose en segundo plano mientras el resto de codigo termina su ejecucion y
	cuando cada uno de estos procesos termine su ejecucion, obtendremos los resultados del mismo
	incluso antes de que se acabe de ejecutar todo nuestro codigo, es similar a la asincronia en JavaScript*/
	go mensaje("Canal 1", Canal1)
	go mensaje("Canal 2", Canal2)

	for i := 0; i < 2; i++ {
		select {
		case chat1 := <-Canal1:
			fmt.Println("Se ha alterado el", chat1)
		case chat2 := <-Canal2:
			fmt.Println("Se ha alterado el", chat2)
		}
	}
}
