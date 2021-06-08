package main

import "fmt"

func main() {
	canal := make(chan string, 200)
	canal <- "Olá mundo"
	canal <- "Programa em GO"
	canal <- "Programa em GO novamente"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
