package main

import "fmt"

/*Los Structs anonimos son Structs a los que no se les define el nombre y a su vez, estos no estan
establecidos fuera de la funcion main sino que estos funcionan dentro de esta y solo seran ejecutados
esa vez*/

func main() {

	Guardado := struct {
		//Parametros
		Nombre   string
		Apellido string
	}{
		//Valores
		"Victor",
		"Ramirez",
	}

	fmt.Println(Guardado) //{Victor Ramirez}

	//Tambien podemos llamar los valores accediendo con DotNotation
	fmt.Println(Guardado.Nombre) //Victor

	fmt.Println("Bienvenido señor", Guardado.Apellido) //Bienvenido señor Ramirez
}
