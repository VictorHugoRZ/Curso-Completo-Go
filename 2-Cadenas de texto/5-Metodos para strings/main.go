// The Go programming language provides a number of methods for working with strings. These methods can be used to perform a variety of tasks, such as:
// * Concatenating strings
// * Converting strings to lowercase or uppercase
// * Searching for a substring within a string
// * Replacing a substring within a string
// * Trimming whitespace from a string
// * Splitting a string into substrings
// The following is a complete list of the string methods in Go:
// * `Append()`: Appends a string to the end of another string.
// * `Cap()`: Returns the capacity of a string.
// * `Char()`: Returns the character at the specified index in a string.
// * `Compare()`: Compares two strings and returns an integer indicating their order.
// * `Contains()`: Returns a boolean indicating whether a substring is contained within a string.
// * `Copy()`: Copies a string to a new string.
// * `Count()`: Returns the number of occurrences of a substring within a string.
// * `Equal()`: Returns a boolean indicating whether two strings are equal.
// * `Fields()`: Splits a string into substrings separated by whitespace.
// * `HasPrefix()`: Returns a boolean indicating whether a string starts with a specified prefix.
// * `HasSuffix()`: Returns a boolean indicating whether a string ends with a specified suffix.
// * `Index()`: Returns the index of the first occurrence of a substring within a string.
// * `LastIndex()`: Returns the index of the last occurrence of a substring within a string.
// * `Len()`: Returns the length of a string.
// * `Lower()`: Converts a string to lowercase.
// * `Map()`: Applies a function to each character in a string.
// * `New()`: Creates a new string with the specified contents.
// * `Repeat()`: Returns a new string that is the specified number of copies of the original string.
// * `Replace()`: Replaces all occurrences of a substring within a string with another substring.
// * `ReplaceAll()`: Replaces all occurrences of a substring within a string with another substring, using a function to determine the replacement string.
// * `Runes()`: Returns a slice of runes that represents the contents of a string.
// * `ToLower()`: Converts a string to lowercase.
// * `ToUpper()`: Converts a string to uppercase.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var cadena string = "ESTE ES UN STRING EN MAYUSCULAS"

	fmt.Println(cadena) //ESTE ES UN STRING EN MAYUSCULAS

	cadena = strings.ToLower(cadena) //Transforma el string que nosotros le indicamos a minusculas.
	fmt.Println(cadena)              //este es un string en mayusculas

	cadena = strings.ToUpper(cadena) //Transforma el string que nosotros le indicamos a mayusculas.
	fmt.Println(cadena)              //ESTE ES UN STRING EN MAYUSCULAS

	cadena = strings.Replace(cadena, "E", "-", 3)
	fmt.Println(cadena) //-ST- -S UN STRING EN MAYUSCULAS
	//Primero le indicamos que strings se modificara "cadena1", Segundo el caracter que queremos reemplazar "E",
	//Tercero el caracter que reemplazara al anterior "-" y Cuarto, el numero de veces que se repetira esta indicacion "3".
	//Tenemos que tener cuidado al indicar si el caracter que se quiere cambiar es una letra mayuscula o minuscula.

	cadena = strings.ReplaceAll(cadena, "S", ".")
	fmt.Println(cadena) //-.T- -. UN .TRING EN MAYU.CULA.
	//A diferencia del metodo anterior, aqui no definimos el numero de veces que se reemplazaran caracteres,
	//esto reemplaza todos los caracteres iguales que encuentre a lo largo del string.
}
