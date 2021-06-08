package main

import (
	"fmt"
	"linha_de_comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Start")

	aplicacao := app.Gerar()
	aplicacao.Run(os.Args)
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
