package main

/*
Inicializamos nuestros paquetes con el comando
go mod init <nombre del paquete nuevo>
*/

import (
	"fmt"

	"./paquete"
)

func main() {
	var datos paquete.Datos
	datos.Nombre = "Victor"
	datos.Apellido = "Ramirez"
	datos.Edad = "24"

	fmt.Println(datos)
}
