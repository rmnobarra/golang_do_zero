package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Structs")

	var usuario1 usuario
	usuario1.nome = "Macaco"
	usuario1.idade = 69
	fmt.Println(usuario1)

	enderecoExemplo := endereco{"Elm Street", 69}

	usuario2 := usuario{"Florzinha", 10, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Lindinha"}
	usuario4 := usuario{idade: 9}

	fmt.Println(usuario3)
	fmt.Println(usuario4)
}
