package main

import "fmt"

// Interface vinculante y canalizadora de metodos (funciones)
type OperacionArit interface {
	resultado() int // Metodo referente a las funciones de las lineas 17 y 27
}

// Struct Suma
type Suma struct {
	numero  int
	numero1 int
}

// Funcion de Metodo referente a el struct Suma
func (s Suma) resultado() int {
	return s.numero + s.numero1
}

// Struct Resta
type Resta struct {
	numero2 int
	numero3 int
}

// Funcion de Metodo referente a el struct Resta
func (r Resta) resultado() int {
	return r.numero2 - r.numero3
}

/*Cuando esta funcion recibe los parametros, al invocarse vincula estos a la interface, la interface
canaliza la operacion hacia el metodo correspondiente de acuerdo al parametro que le hayamos pasado
originalmente a esta funcion y el metodo es ejecutado, todo esto para retornarnos la operacion
correspondiente*/
func Operacion(o OperacionArit) {
	fmt.Println("El resultado de la operacion es:", o.resultado())
}

func main() {
	var decision string
	suma := Suma{}
	resta := Resta{}

	fmt.Println("Calculadora")
	fmt.Println("¿Desea realizar una suma o resta?")
	fmt.Scanln(&decision)
	if decision == "suma" || decision == "Suma" || decision == "SUMA" {
		fmt.Println("Perfecto, ingrese los valores de la suma:")
		//fmt.Println("Primer valor; ")
		fmt.Scanln(&suma.numero)
		//fmt.Println("Segundo valor;")
		fmt.Scanln(&suma.numero1)
		Operacion(suma)
	} else if decision == "resta" || decision == "Resta" || decision == "RESTA" {
		fmt.Println("Perfecto, ingrese los valores de la resta:")
		//fmt.Println("Primer valor; ")
		fmt.Scanln(&resta.numero2)
		//fmt.Println("Segundo valor;")
		fmt.Scanln(&resta.numero3)
		Operacion(resta)
	} else {
		fmt.Println("Ingrese una operacion valida para realizar")
	}

}
