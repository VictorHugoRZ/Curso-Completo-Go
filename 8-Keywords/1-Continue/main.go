package main

import "fmt"

func main() {
	//Continue
	//La palabra clave continue nos ayuda a omitir cierta parte de nuestro codigo

	for carga := 0; carga <= 10; carga++ {
		if carga == 4 || carga == 5 || carga == 6 {
			continue
		}
		fmt.Println(carga)
	}
	//Se omiten los numero 4, 5, 6 al utilizar el continue
	//1
	//2
	//3
	//7
	//8
	//9
	//10
}
