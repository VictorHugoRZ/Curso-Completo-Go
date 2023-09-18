package main

import "fmt"

type ListaInvitados struct {
	Numero   int8
	Nombre   string
	Apellido string
	Edad     int8
}

func main() {
	invitado1 := ListaInvitados{001, "Victor", "Ramirez", 25} //Se utilizan llaves {}
	fmt.Println(invitado1)                                    //{1 Victor Ramirez 25}

	var invitado2 = ListaInvitados{002, "Hugo", "Zamora", 25} //Podemos declararlos como variables
	fmt.Println(invitado2)                                    //{2 Hugo Zamora 25}

	/*Tambien podemos acceder a la informacion de un Struct en manera de objeto con Dot Notation*/
	fmt.Println("Invitado:", invitado1.Numero, "Nombre:", invitado1.Nombre, "Edad:", invitado1.Edad)
	//Invitado: 1 Nombre: Victor Edad: 25

	fmt.Println("Invitado:", invitado2.Numero, "Nombre:", invitado2.Nombre, "Edad:", invitado2.Edad)
	//Invitado: 2 Nombre: Hugo Edad: 25
}
