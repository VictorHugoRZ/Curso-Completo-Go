package main

import "fmt"

func main() {
	bolsa := []string{}

	bolsa = append(bolsa, "Una bolsa de Pan")
	bolsa = append(bolsa, "Dos libras de Carne")
	bolsa = append(bolsa, "6 libras de Arroz")
	bolsa = append(bolsa, "2 bolsas de Verdura")
	bolsa = append(bolsa, "3 galones de Agua")

	fmt.Println("Maria ha comprado los siguientes productos:", bolsa)
}
