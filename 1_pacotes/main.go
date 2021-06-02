package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("hello again sucker")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("rmnobarra.gmail.com")
	fmt.Println(erro)
}
