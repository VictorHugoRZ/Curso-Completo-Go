package main

import "fmt"

// Crea 4 variables las cuales 1 será decimal y las otras 3 enteras,
// la primera variable se restara con la segunda, el resultado de la resta
// se multiplicara con la tercer variable y se dividirá entre la cuarta variable.

func main() {
	var num1 int = 5
	var num2 int = 5
	var num3 int = 10
	var num4 int = 2

	var resta int = num1 - num2 //Operacion resta realizada
	fmt.Println(resta)

	var multiplicacion = resta * num3 //Operacion multiplicacion realizada
	fmt.Println(multiplicacion)

	var division = multiplicacion / num4 //Operacion division realizada
	fmt.Println(division)

}
