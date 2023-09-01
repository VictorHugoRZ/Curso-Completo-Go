package main

import "fmt"

func main() {
	var letra string
	fmt.Println("----¿Es vocal?----")
	fmt.Println("Ingresa una letra:")
	fmt.Scanln(&letra)

	if letra == "a" || letra == "A" {
		fmt.Println("La vocal es A")
	} else if letra == "e" || letra == "E" {
		fmt.Println("La vocal es E")
	} else if letra == "i" || letra == "I" {
		fmt.Println("La vocal es I")
	} else if letra == "o" || letra == "O" {
		fmt.Println("La vocal es O")
	} else if letra == "u" || letra == "U" {
		fmt.Println("La vocal es U")
	} else {
		fmt.Println("La letra ingresada no es una vocal")
	}
}
