package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	var array1 [5]string
	array1[0] = "posição 1"

	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}

	fmt.Println(array2)

	slice := []int{10, 11, 12, 13, 14}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))

	slice = append(slice, 69)

	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

}
