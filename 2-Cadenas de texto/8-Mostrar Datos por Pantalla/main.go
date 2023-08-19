package main

import "fmt"

func main() {
	saludo, despedida := "Hola", "Adios"

	fmt.Println(saludo, "Victor")
	fmt.Println(despedida, "Hugo")

	nombre := "Victor"
	edad := 24

	//Printf (format) no agrega saltos de linea, podemos agregar \n al final del formato para agregar el salto de linea.
	fmt.Printf("Hola %s, tu edad es: %d\n", nombre, edad) //%s se usa para strings y %d para numeros, a estos se les llaman verbos en Go.

	fmt.Printf("Hola %v, tu edad es: %v\n", nombre, edad) //%v se usa cuando no sabemos el tipo de dato de las variables que usaremos en el formato.

	var mensaje string

	fmt.Println("Ingresa un mensaje:")
	fmt.Scanln(&mensaje)
	fmt.Printf("Tu mensaje es: %s", mensaje)

}
