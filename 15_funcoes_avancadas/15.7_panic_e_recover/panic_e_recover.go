package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar a execução")
	}

}
func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("a média é exatamante 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("pós execução")

}
