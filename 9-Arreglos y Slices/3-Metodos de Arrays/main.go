package main

import "fmt"

func main() {
	array := [3]int8{1, 2, 3}
	fmt.Println(array) //[1 2 3]

	/*Para conocer la longitud total de un arreglo podemos utilizar el metodo len, recordemos que esto nos
	dara el numero total de elementos que el array podra almacenar, aunque alguno de sus espacios aun no este
	ocupado, NO nos dara el numero de elementos actuales en el array, este ultimo valor puede variar
	y podria ser inexacto*/

	fmt.Println(len(array)) //3

	array2 := [5]int8{1, 2, 3, 4}
	fmt.Println(array2)      //[1 2 3 4 0]
	fmt.Println(len(array2)) //5
	fmt.Println(cap(array2)) //5
	//Cap nos permite saber la capacidad del array

	/*Here is a list of all array methods in Go with a brief explanation of each one:

	- append(): Adds one or more elements to the end of an array and returns the new array.

	- cap(): Returns the capacity of an array, which is the maximum number of elements it can hold.

	- copy(): Copies the elements of an array to another array.

	- len(): Returns the length of an array, which is the number of elements it contains.

	- make(): Creates a new array with the specified length and capacity.

	- delete(): Deletes an element from an array.

	- insert(): Inserts an element into an array at a specified index.

	- pop(): Removes and returns the last element of an array.

	- push(): Adds one or more elements to the end of an array.

	- resize(): Changes the length and capacity of an array.

	- reverse(): Reverses the order of the elements in an array.

	- search(): Searches for an element in an array and returns its index.

	- sort(): Sorts the elements of an array in ascending or descending order.

	- swap(): Swaps the values of two elements in an array.

	- truncate(): Removes all elements from an array except for the first n elements.

	- clear(): Removes all elements from an array.

	- empty(): Returns true if the array is empty, and false otherwise.

	- full(): Returns true if the array is full, and false otherwise.

	- set(): Sets the value of an element in an array at a specified index.

	- get(): Returns the value of an element in an array at a specified index.

	- slice(): Returns a slice of an array.

	- join(): Joins an array of strings into a single string.

	- split(): Splits a string into an array of strings.

	- reverseSlice(): Reverses the order of the elements in a slice.

	- copySlice(): Copies the elements of a slice to another slice.

	- appendSlice(): Adds one or more elements to the end of a slice and returns the new slice.

	- prependSlice(): Adds one or more elements to the beginning of a slice and returns the new slice.

	- insertSlice(): Inserts an element into a slice at a specified index and returns the new slice.

	- deleteSlice(): Deletes an element from a slice */
}
