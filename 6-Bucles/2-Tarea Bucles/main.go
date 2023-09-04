package main

import "fmt"

func main() {
	fmt.Println("¿Cual es el ciclo de la vida?")
	var cicloDeVida string

	for i := 0; i <= 4; i++ {

		fmt.Println(cicloDeVida)
		if i == 0 {
			cicloDeVida = "Naces"
		} else if i == 1 {
			cicloDeVida = "Creces"
		} else if i == 2 {
			cicloDeVida = "Te reproduces"
		} else if i == 3 {
			cicloDeVida = "Mueres"
		}
	}
}
