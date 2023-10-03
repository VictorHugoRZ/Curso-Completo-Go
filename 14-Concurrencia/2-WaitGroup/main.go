package main

import (
	"fmt"
	"sync"
)

func MuchoTexto(Palabra string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(Palabra)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hola a todos!")

	wg.Add(1)

	go MuchoTexto("Uso de WaitGroup", &wg)

	wg.Wait()
}
