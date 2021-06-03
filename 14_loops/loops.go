package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	fmt.Println("-----")

	i := 0

	for i < 10 {
		i++
		fmt.Println("incrementando i", i)
		//time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("incrementando j", j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"florzinha", "lindinha", "docinho"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "leonardo",
		"sobrenome": "silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//loop infinito
	for {
		fmt.Println("executando para todo o sempre")
		time.Sleep(time.Second)
	}
}
