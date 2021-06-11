package main

import (
	"fmt"
	"introducao_testes/enderecos"
)

func main() {
	TipoDeEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(TipoDeEndereco)
}
