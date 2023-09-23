package main

import "fmt"

/*Podemos definir una clase de Herencia dentro de los Structs, ya que podemos pasar las propiedades
de un Struct a otro embebiendolos, esto quiere decir que se unen las propiedades de el Struct Padre
al Struct Hijo*/

type perro struct {
	raza      string
	tamaño    string
	saludable bool
	peso      int8
}

type pug struct {
	perro
	nombre   string
	apodo    string
	vacunado bool
}

func main() {

	var perro1 = pug{
		perro{
			"Pug",
			"Pequeño",
			true,
			15,
		},
		"Porfirio",
		"gordito",
		true,
	}

	fmt.Println(perro1) //{{Pug Pequeño true 15} Porfirio gordito true}

	fmt.Println(perro1.nombre) //Porfirio

}
