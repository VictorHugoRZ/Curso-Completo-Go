package main

import (
	"fmt"
	"time"
)

func MuchoTexto(Palabra string) {
	fmt.Println(Palabra)
}

func main() {

	MuchoTexto("Hola a todos!")
	go MuchoTexto("Estas son las Go Routines")

	time.Sleep(time.Second * 1)

}
