package main

import "fmt"

/*Los structs son estructuras de datos parecidas a los objetos, estas nos permiten definir modelos
u objetos que tendran una estructura de datos definida con sus tipos de datos cada una, podemos usarlos
para definir posterirmente instancias de multiples cosas u objetos.*/

type Carro struct {
	Marca string
	Año   int64
	Color string
}

func main() {
	miCarro := Carro{"Mazda", 2013, "Rojo"}
	fmt.Println(miCarro) //{Mazda 2013 Rojo}

	miCarro2 := Carro{"Mercedez Benz", 2015, "Plata"}
	fmt.Println(miCarro2) //{Mercedez Benz 2015 Plata}
}
