package main

import "fmt"

//Todos tenemos una rutina diaria en nuestro dia a dia, ¿como representarias la tuya con un bucle?.

func main() {
	var hora int8
	var actividad string

	for i := 0; i < 24; i++ {
		hora++
		fmt.Println("Son las", hora, "hrs")
		if hora == 7 {
			actividad = "Hora de levantarse"
			fmt.Println(actividad)
		}
		if hora == 10 {
			actividad = "A limpiarle a los perritos"
			fmt.Println(actividad)
		}
		if hora == 11 {
			actividad = "A desayunar"
			fmt.Println(actividad)
		}
		if hora == 15 {
			actividad = "A de comer"
			fmt.Println(actividad)
		}
		if hora == 17 {
			actividad = "Hora de estudiar"
			fmt.Println(actividad)
		}
		if hora == 21 {
			actividad = "A de cenar"
			fmt.Println(actividad)
		}
		if hora == 23 {
			actividad = "Hora de dormir"
			fmt.Println(actividad)
		}
	}
}
