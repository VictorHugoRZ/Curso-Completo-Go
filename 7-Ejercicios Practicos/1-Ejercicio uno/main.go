package main

import "fmt"

//Sofia tiene visitas en su casa, un total de 20 personas. Ella debe atender a sus invitados y cada
//uno de ellos le ha pedido un vaso de agua, a continuacion Sofia va a su cocina y comienza a llenar
//un vaso tras otro hasta completar los 20 vasos para sus invitados.
//Utiliza un bucle for para representar y automatizar esta situacion.

func main() {
	var vaso int8
	var invitado int8

	for i := 0; i < 20; i++ {
		vaso++
		invitado++
		fmt.Println("Sofia ha tomado el vaso numero", vaso, "y lo llena de agua.")
		fmt.Println("Ahora se lo entrega al invitado numero", invitado)
	}
}
