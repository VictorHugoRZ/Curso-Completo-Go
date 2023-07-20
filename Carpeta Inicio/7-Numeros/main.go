package main

import (
	"fmt"
	"reflect"
)

// En Go podemos asignar variables y además determinar la cantidad especifica de memoria que pretendemos
// utilizar, esto en cuanto a entornos de ejecución, refiriéndonos a el rango que definiremos para las
// variables y así limitar el uso de memoria.

// Tabla de numeros y rangos de bits en variables numericas en la carpeta 7, hacer click derecho y
// abrir desde la carpeta.

func main() {
	var numero int = 50

	fmt.Println(numero)                 //50
	fmt.Println(reflect.TypeOf(numero)) //int

	numeroFlotante := 20.50

	fmt.Println(numeroFlotante)                 //20.50
	fmt.Println(reflect.TypeOf(numeroFlotante)) //float64

	var num int8 = 127 //El valor de este entero es de 8 bits y su rango está entre -128 y 127, no podemos rebasar ese limite o estariamos provocando un desbordamiento de memoria.

	fmt.Println(num)                 //127
	fmt.Println(reflect.TypeOf(num)) //int8

	var num2 int16 = 32767 //Esta variable tiene un valor de 16 bits y su rango esta entre -2^15 y 2^15-1, es decir: -32,768 y 32,767, no podemos rebasar ese limite o estariamos provocando un desbordamiento de memoria.

	fmt.Println(num2)                 //32767
	fmt.Println(reflect.TypeOf(num2)) //int16

	var num3 uint8 = 255 //Este tipo de dato no permite valores negativos, solo valores a partir de 0 a 255.

	fmt.Println(num3)                 //255
	fmt.Println(reflect.TypeOf(num3)) //uint8

	var num4 float32 = 23.98 //Para los numeros flotantes solo exxisten los tipos de datos float32 y float64

	fmt.Println(num4)                 //23.98
	fmt.Println(reflect.TypeOf(num4)) //float32

	//Las medidas de rango de los numeros flotantes son las mismas a los enteros de los bits correspondientes

	//TODO ESTO NOS AYUDA PARA PODER LIMITAR LA MEMORIA, NOS SERA UTIL EN LA CREACION DE PROYECTOS GRANDES
}
