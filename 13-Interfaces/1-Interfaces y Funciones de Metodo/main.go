package main

import "fmt"

type Persona struct { //Struct referenciado en la funcion de metodo de la linea 17
	Nombre string
}

type Institucion struct {
	Nombre string
}

/*Funciones de metodo: estas funciones pueden hacer referencia a un struct agregando un parentesis
antes del nombre que les asignamos*/

// Referencia
func (p Persona) mostrar() string {
	return p.Nombre
}

func (i Institucion) mostrar() string {
	return i.Nombre
}

/*Go no nos lanza un error al nombrar a las funciones anteriores de la misma manera ya que no son simples
funciones, sino metodos que hacen referencia a Structs distintos y para saber a cual de estas llamaremos
haremos una interfaz*/

type Datos interface {
	mostrar() string
}

/*La interface utiliza el metodo mostrar(), nosotros sabemos que este puede tener dos posibles utilidades,
estas serian utilizar el metodo de la linea 17 o el metodo de la linea 21, nosotros solo llamamos al metodo
mostrar() y Go utilizara el metodo pertinente dependiendo de a que struct hagamos referencia*/

/*Ahora la funcion Peticion() de la linea 41 hace referencia a la interface de la linea 29 y detecta que
nosotros al usar la interface podemos acceder a todos los metodos que esta tenga dentro de si, podemos
acceder a estos utilizando DotNotation*/

func Peticion(d Datos) { //-----------------DonNotation
	fmt.Println("Los datos ingresados son:", d.mostrar())
}

/*Ahora creamos nuestras variables que le dan valores a cada uno de los structs que definimos al inicio de
nuestro paquete, al llamar a nuestra funcion Peticion() de la linea 41 y pasarle cualquiera de nuestras
variables: Nombre de la linea 52 o Escuela de la linea 56, la funcion Peticion() detectara a que struct nos
estamos refiriendo y con base en eso usara el metodo mostrar de la interfaz que haga referencia a el mismo
struct de la variable que le pasemos como parametro*/

func main() {

	var Nombre = Persona{
		"Victor",
	}

	var Escuela = Institucion{
		"CONAMAT",
	}

	Peticion(Escuela) //Los datos ingresados son: CONAMAT
	Peticion(Nombre)  //Los datos ingresados son: Victor

}
