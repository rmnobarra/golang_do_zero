package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero1 int16 = 1000

	var numero2 uint64 = 69

	//int32
	var numero3 rune = 6969

	//uint8
	var numero4 byte = 123

	fmt.Println(numero1, numero2, numero3, numero4)

	var numero5 float32 = 69.69
	fmt.Println(numero5)

	var numero6 float64 = 121324.98
	fmt.Println(numero6)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var texto2 int
	fmt.Println(texto2)

	var boolean1 bool = true
	fmt.Println(boolean1)

	var boolean2 bool
	fmt.Println(boolean2)

	var erro error = errors.New("Erro do errors")
	fmt.Println(erro)
}
