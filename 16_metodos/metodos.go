package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Salvando os dados do %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"usuario 1", 69}
	usuario1.salvar()

	usuario2 := usuario{"Usuario 2", 18}
	usuario2.salvar()

	maiordeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiordeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
