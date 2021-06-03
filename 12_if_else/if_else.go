package main

import "fmt"

func main() {
	fmt.Println("if else")
	fmt.Println("-------")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero maior que zero")
	} else {
		fmt.Println("Numero menor do que zero")
	}

}
