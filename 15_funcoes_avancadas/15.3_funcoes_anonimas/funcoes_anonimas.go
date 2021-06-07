package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Hello anonymous")

	fmt.Println(retorno)
}
