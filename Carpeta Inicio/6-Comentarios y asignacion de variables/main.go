package main

import "fmt" //Asi asignamos comentarios

//Podemos asignar variables fuera de las funciones

var numero int //Cuando no asignamos valor a las variables de tipo entero se setean automaticamente en cero 0

var texto string = "Victor" //Podemos setear las variables con un valor definido

func main() {
	fmt.Println(numero) //0
	fmt.Println(texto)  //Victor

	var texto2 string   //No asignar valor a variables de tipo string setea la variable en un string vacio
	fmt.Println(texto2) //""

	var numero2 int //El valor es cero 0, ya que no asignamos valor pero esta vez se declaro dentro de la funcion
	fmt.Println(numero2)

	variable := 'A'       //Declaracion de variables, si se usan comillas simples es para un solo caracter 'A'
	fmt.Println(variable) //65, esto por que el valor se traduce en codigo asccii

	variable2 := "Hola"    //Si se usan comillas dobles es para un string
	fmt.Println(variable2) //Hola

	//Podemos reasignar estas variables de esta manera

	variable = 'B'        //Se reasigno el valor de la variable de la linea 21
	fmt.Println(variable) //66 por el codigo asccii

	variable2 = "Adios"    //Se reasigno el valor de la variable de la linea 21
	fmt.Println(variable2) //Adios

	//Creacion y asignacion multiple de variables

	var v1, v2, v3 int //Creacion de multiples variables al mismo tiempo sin asignacion de valores
	fmt.Println(v1)    //0
	fmt.Println(v2)    //0
	fmt.Println(v3)    //0

	var v4, v5, v6 int = 12, 23, 34 //Creacion de multiples variables al mismo tiempo con asignacion de valores
	fmt.Println(v4)                 //12
	fmt.Println(v5)                 //23
	fmt.Println(v6)                 //34

	//Asignacion multiple de variables sIn delaracion con palabra reservada

	variableA, variableB, variableC := 10, 20, 30
	fmt.Println(variableA) //10
	fmt.Println(variableB) //20
	fmt.Println(variableC) //30
}
